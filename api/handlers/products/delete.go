package products

import (
	"FinalProject/models"
	"FinalProject/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *ProductHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		currentUser, ok := ctx.Get("currentUserId")
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "You need to login to get detail"})
			return
		}
		idProduct := ctx.Param("id")

		product := models.Product{}
		h.Db.Find(&product, idProduct)

		permission := utils.CheckPermission(currentUser.(uint), product.UserID)

		if permission {

			h.Db.Delete(&product)

			ctx.JSON(http.StatusAccepted, gin.H{"message": fmt.Sprintf("Product with id %v is deleted", idProduct)})
			return
		}
		ctx.JSON(http.StatusForbidden, gin.H{"message": "You have no permission "})
	}
}
