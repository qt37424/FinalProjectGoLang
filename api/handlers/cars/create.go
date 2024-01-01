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
		Product models.Product
	}
)

func (h *ProductHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// jwtToken, err := ctx.Cookie("Authorization")
		// if err != nil {
		// 	ctx.AbortWithStatus(http.StatusUnauthorized)
		// }
		creater, ok := ctx.Get("currentUserId")
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Cannot get user Id"})
			return
		}

		body := &createBody{}
		if err := ctx.BindJSON(body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		product := models.Product{
			Name:   body.Name,
			Price:  body.Price,
			UserID: creater.(uint),
		}
		result := h.Db.Create(&product)
		if result.Error != nil {
			// This error cause from DB so server cannot catch it
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		ctx.JSON(http.StatusOK, &createRes{Product: product})
	}
}
