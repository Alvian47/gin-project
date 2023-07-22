package middleware

import (
	"diary_api/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := helper.ValidateJWT(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": gin.H{
					"errorValidateJWT": "Authentication required",
					"msg":              err.Error(),
				},
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
