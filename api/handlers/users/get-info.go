package users

import (
	"FinalProject/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) GetInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users := &[]models.User{}
		h.Db.Find(users)
		ctx.JSON(http.StatusOK, gin.H{"data": users})
	}
}
