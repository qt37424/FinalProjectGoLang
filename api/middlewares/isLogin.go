package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsLogined() gin.HandlerFunc { // is not done
	return func(ctx *gin.Context) {
		_, ok := ctx.Get("currentUserId")
		if ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You have been logined!"})
			return
		}
		ctx.Next()
	}
}
