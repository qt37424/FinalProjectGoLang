package products

import (
	"FinalProject/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	getDetailPathParams struct {
		ID uint `uri:"id" binding:"required"`
	}
	getDetailRes struct {
		ID    uint             `json:"id"`
		Name  string           `json:"name"`
		Price uint             `json:"price"`
		User  getDetailUserRes `json:"user"`
	}
	getDetailUserRes struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
	}
)

func (h *ProductHandler) GetDetail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		currentUser, ok := ctx.Get("currentUserId")
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "You need to login to get detail"})
			return
		}
		pathParms := getDetailPathParams{}

		if err := ctx.ShouldBindUri(&pathParms); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		product := models.Product{}
		result := h.Db.Preload("User").First(&product, pathParms)

		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Data with id not found"})
				return
			}
			ctx.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
			return
		}

		res := getDetailRes{ // ta có thể sử dụng struct của get list để làm
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price,
			User: getDetailUserRes{
				ID:       product.User.ID,
				Username: product.User.Username,
			},
		}
		if currentUser == res.User.ID {
			ctx.JSON(http.StatusOK, gin.H{"data": res})
		} else {
			ctx.JSON(http.StatusForbidden, gin.H{"message": "You have no permission!"})
		}

		// // Cách 1
		// product := models.Product{}
		// id := ctx.Param("id")
		// h.Db.Find(&product, id)

		// data := getListProductRes{
		// 	ID:    product.ID,
		// 	Name:  product.Name,
		// 	Price: product.Price,
		// }
		// ctx.JSON(http.StatusOK, gin.H{"data": data})
	}
}
