package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (*ProductHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(http.StatusOK, gin.H{"message": "updated info of id: " + id})
	}
}
