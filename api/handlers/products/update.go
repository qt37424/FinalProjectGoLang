package products

import (
	"FinalProject/models"
	"FinalProject/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *ProductHandler) Update() gin.HandlerFunc {
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
			body := &createBody{}
			if err := ctx.BindJSON(body); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			h.Db.Model(&product).Updates(models.Product{Name: body.Name, Price: body.Price})
			data := getListProductRes{
				ID:    product.ID,
				Name:  product.Name,
				Price: product.Price,
			}
			ctx.JSON(http.StatusAccepted, gin.H{"data": data})
			return
		}
		ctx.JSON(http.StatusForbidden, gin.H{"message": "You have no permission "})
	}
}
