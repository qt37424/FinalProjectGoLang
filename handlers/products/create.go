package products

import (
	"FinalProject/models"
	"net/http"

	"FinalProject/utils"

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
		creater, err := utils.ExtractUsernameFromToken(ctx.Request)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		body := &createBody{}
		if err := ctx.ShouldBindJSON(body); err != nil { // here we also have a method call MustBindJSON
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		product := &models.Product{
			Name:   body.Name,
			Price:  body.Price,
			UserID: creater,
		}
		result := h.Db.Create(product)
		if result.Error != nil {
			// This error cause from DB so server cannot catch it
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		ctx.JSON(http.StatusOK, &createRes{InsertedId: product.ID})
	}
}
