package routers

import (
	"FinalProject/handlers/products"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) AddProductRouter(apiRouter *gin.RouterGroup) {
	productRouter := apiRouter.Group("products")
	handler := products.ProductHandler{Db: r.DB}

	productRouter.GET("", handler.GetAll())

	productRouter.GET("/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(http.StatusOK, gin.H{"message": "get by id: " + id})
	})

	productRouter.POST("")
	productRouter.PUT("/:id")
	productRouter.DELETE("/:id")
}
