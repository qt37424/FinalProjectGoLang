package products

import (
	"FinalProject/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	createBody struct {
		Name  string `json:"name" binding:"required"`
		Price uint   `json:"price" binding:"required,min=100"`
	}
	createRes struct {
		InsertedId uint `json:"insertId"`
	}
)

func (h *ProductHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := &createBody{}
		if err := ctx.ShouldBindJSON(body); err != nil { // here we also have a method call MustBindJSON
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		product := &models.Product{
			Name:   body.Name,
			Price:  body.Price,
			UserID: 1, // hardcode to study when do project re-config
		}
		result := h.Db.Create(product)
		if result.Error != nil {
			// Lỗi này từ phía DB nên sẽ không catch được
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		ctx.JSON(http.StatusOK, &createRes{InsertedId: product.ID})
	}
}
