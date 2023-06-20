package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (*ProductHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "get list"})
	}
}
