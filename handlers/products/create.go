package products

import (
	"FinalProject/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateBody struct {
	Name  string `json:"name" binding:"required"`
	Price uint   `json:"price" binding:"required,min=100"`
}

func (h *ProductHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := &CreateBody{}
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
		ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Data have been created with ID %v", product.ID)})
	}
}
