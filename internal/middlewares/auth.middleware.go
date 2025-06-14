package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/vuvanthanhtb/go-ecommerce-backend/packages/response"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token != "valiad-token" {
			response.ErrorResponse(ctx, response.ErrCodeParamInvalid)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
