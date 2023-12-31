package products

import (
	"FinalProject/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	getListRes struct {
		Data []getListProductRes `json:"data"`
	}
	getListProductRes struct { // I think we should pass the field for token id when sending data
		ID    uuid.UUID `json:"id"`
		Name  string    `json:"name"`
		Price uint      `json:"price"`
	}
	RepositoryResult struct {
		Result interface{}
		Error  error
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
