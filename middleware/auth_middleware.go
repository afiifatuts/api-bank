package middleware

import (
	"net/http"

	"github.com/afiifatuts/bankmnc/helper"
	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Missing authorization token",
			})
			ctx.Abort()
			return
		}

		payload, err := helper.VerifyToken(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Invalid or expired token",
			})
			ctx.Abort()
			return
		}

		// Token is valid, proceed to the next handler
		ctx.Set("payload", payload)
		ctx.Next()
	}
}
