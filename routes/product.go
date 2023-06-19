package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) AddProductRouter(apiRouter *gin.RouterGroup) {
	productRouter := apiRouter.Group("products")
	productRouter.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "get list"})
	})

	productRouter.GET("/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(http.StatusOK, gin.H{"message": "get by id: " + id})
	})

	productRouter.POST("")
	productRouter.PUT("/:id")
	productRouter.DELETE("/:id")
}
