package api

import (
	"net/http"

	"github.com/Mhoanganh207/BankBE/database"
	"github.com/Mhoanganh207/BankBE/models"
	"github.com/gin-gonic/gin"
)

type AccountRequest struct {
	Balance     int64  `json:"balance"`
	Currency    string `json:"currency"`
	CountryCode int    `json:"country_code"`
}

func (r Routes) addAccountsRoute(server *Server) {
	accounts := r.router.Group("/api/accounts")
	accounts.Use(authMiddleware(server.tokenService))
	accounts.GET("", server.getAccount)
	accounts.POST("", server.createAccount)
}

func (s *Server) createAccount(ctx *gin.Context) {
	var req AccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	subject, _ := ctx.Get("subject")
	owner := s.tokenService.GetSubject(subject.(string))

	account := &models.Account{
		Owner:       owner,
		Balance:     req.Balance,
		Currency:    req.Currency,
		CountryCode: req.CountryCode,
	}

	if err := database.CreateAccount(account, s.db); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, account)
}

func (s *Server) getAccount(ctx *gin.Context) {
	subject, _ := ctx.Get("subject")
	owner := s.tokenService.GetSubject(subject.(string))
	account, err := database.GetAccount(owner, s.db)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, account)
}
