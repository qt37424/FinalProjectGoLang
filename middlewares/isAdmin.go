package middlewares

import "github.com/gin-gonic/gin"

func CheckIsAdmin() gin.HandlerFunc { // Not implement
	return func(ctx *gin.Context) {

		ctx.Next()
	}
}
