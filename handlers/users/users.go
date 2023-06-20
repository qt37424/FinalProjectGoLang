package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (*UserHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "get one user"})
	}
}
