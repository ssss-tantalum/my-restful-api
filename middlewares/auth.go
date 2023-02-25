package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ssss.tantalum/my-restful-api/auth"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "request does not contain an access token"})
			ctx.Abort()
			return
		}
		err := auth.ValidateToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
		}
		ctx.Next()
	}
}
