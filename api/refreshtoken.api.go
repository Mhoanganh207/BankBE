package api

import (
	"net/http"
	"time"

	"github.com/Mhoanganh207/BankBE/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type RefreshTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

func (r Routes) addTokenRoute(server *Server) {
	tokens := r.router.Group("/api/token")
	tokens.POST("/refresh", server.refreshToken)
}

func (s *Server) refreshToken(ctx *gin.Context) {
	var req RefreshTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	refreshPayload, err := s.tokenService.ValidateToken(req.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	payload, err := refreshPayload.GetSubject()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	session, err := database.GetSession(uuid.MustParse(s.tokenService.GetSubject(payload)), s.db)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	accessToken, _, err := s.tokenService.GenerateToken(session.Username, s.config.Duration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	expiresAt := time.Now().Add(s.config.Duration)
	ctx.JSON(http.StatusOK, RefreshTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: expiresAt,
	})
}
