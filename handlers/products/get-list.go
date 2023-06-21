package products

import (
	"FinalProject/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	getListRes struct {
		Data []getListProductRes `json:"data"`
	}
	getListProductRes struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Price uint   `json:"price"`
	}
)

func (h *ProductHandler) GetList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products := []models.Product{}
		h.Db.Find(&products)

		data := []getListProductRes{}
		for _, product := range products {
			data = append(data, getListProductRes{
				ID:    product.ID,
				Name:  product.Name,
				Price: product.Price,
			})
		}
		ctx.JSON(http.StatusOK, getListRes{Data: data})
	}
}
