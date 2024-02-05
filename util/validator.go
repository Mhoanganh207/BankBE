package util

import (
	"net/http"

	"github.com/Mhoanganh207/BankBE/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidCurrency(tf models.Transfer, ctx *gin.Context) bool {
	err := validator.New().Struct(tf)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return false
	}
	return true
}
