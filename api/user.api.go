package api

import (
	"net/http"
	"time"

	"github.com/Mhoanganh207/BankBE/database"
	"github.com/Mhoanganh207/BankBE/models"
	"github.com/Mhoanganh207/BankBE/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type CreateUserResponse struct {
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

type LoginUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserResponse struct {
	SessionID             string    `json:"session_id"`
	AccessToken           string    `json:"access_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshToken          string    `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
}

func (r Routes) addUserRoute(server *Server) {
	users := r.router.Group("/api/users")
	users.POST("", server.createUser)
	users.POST("/login", server.loginUser)
}

func (s *Server) createUser(ctx *gin.Context) {
	var req CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := &models.User{
		Username: req.Username,
		Password: hashedPassword,
		Fullname: req.Fullname,
		Email:    req.Email,
	}

	if err := database.CreateUser(user, s.db); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, CreateUserResponse{
		Username: user.Username,
		Fullname: user.Fullname,
		Email:    user.Email,
	})
}

func (s *Server) loginUser(ctx *gin.Context) {
	var req LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := database.GetUser(req.Username, s.db)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	if !util.CheckPasswordHash(req.Password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	accesstoken, accessPayload, err := s.tokenService.GenerateToken(req.Username, s.config.Duration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	sessionId := uuid.New()
	refreshToken, refreshPayload, err := s.tokenService.GenerateToken(sessionId.String(), s.config.RefreshDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	session := models.Session{
		ID:           sessionId,
		Username:     req.Username,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIP:     ctx.ClientIP(),
		RefreshToken: refreshToken,
		ExpiresAt:    refreshPayload.ExpiresAt.Time,
	}

	if err := database.CreateSession(session, s.db); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := LoginUserResponse{
		SessionID:             session.ID.String(),
		AccessToken:           accesstoken,
		AccessTokenExpiresAt:  accessPayload.ExpiresAt.Time,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiresAt.Time,
	}

	ctx.JSON(http.StatusOK, res)
}
