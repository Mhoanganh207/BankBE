package api

import (
	"fmt"
	"net/http"

	"github.com/Mhoanganh207/BankBE/util"
	"github.com/gin-gonic/gin"
)

func authMiddleware(tokenGenarator util.Generator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			ctx.Abort()
			return
		}
		token = token[7:]
		payload, err := tokenGenarator.ValidateToken(token)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(401, gin.H{"error": "Token is invalid"})
			ctx.Abort()
			return
		}
		subject, err := payload.GetSubject()
		if err != nil {
			ctx.JSON(401, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		ctx.Set("subject", subject)
		ctx.Next()
	}
}
