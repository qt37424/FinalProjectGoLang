package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (*ProductHandler) GetList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "get List Products"})
	}
}
