package api

import (
	"errors"
	"net/http"

	"github.com/Mhoanganh207/BankBE/database"
	"github.com/Mhoanganh207/BankBE/models"
	"github.com/Mhoanganh207/BankBE/util"
	"github.com/gin-gonic/gin"
)

type TransferRequest struct {
	FromAccountId int    `json:"from_account_id" binding:"required,min=1"`
	ToAccountId   int    `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency"`
}

func (r Routes) addTransfersRoute(server *Server) {
	transfers := r.router.Group("/api/transfers")
	transfers.Use(authMiddleware(server.tokenService)).POST("", server.createTransfer)
}

func (s *Server) createTransfer(ctx *gin.Context) {
	var req TransferRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !util.ValidCurrency(models.Transfer{Currency: req.Currency}, ctx) {
		return
	}

	fromAccount, valid := s.validAccount(req.FromAccountId, ctx)
	if !valid {
		return
	}
	subject, _ := ctx.Get("subject")
	owner := s.tokenService.GetSubject(subject.(string))
	if fromAccount.Owner != owner {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": errors.New("YOU_ARE_NOT_THE_OWNER_OF_THIS_ACCOUNT")})
		return
	}

	_, valid = s.validAccount(req.ToAccountId, ctx)
	if !valid {
		return
	}

	transfer := models.Transfer{
		FromAccountId: req.FromAccountId,
		ToAccountId:   req.ToAccountId,
		Amount:        req.Amount,
		Currency:      req.Currency,
	}
	err := database.CreateTransfer(transfer, s.db)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, transfer)
}

func (s *Server) validAccount(id int, ctx *gin.Context) (models.Account, bool) {
	account, err := database.GetAccountById(id, s.db)
	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, err.Error())
			return account, false
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return account, false
	}
	return account, true
}
