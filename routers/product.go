package routers

import (
	handlers "FinalProject/handlers/products"

	"github.com/gin-gonic/gin"
)

func (r *Router) AddProductRouter(apiRouter *gin.RouterGroup) {
	productRouter := apiRouter.Group("products")
	// productRouter.Use(middlewareCheckAuthen)
	handler := handlers.ProductHandler{Db: r.DB}

	productRouter.GET("", handler.GetList())
	productRouter.GET("/:id", handler.GetDetail())
	productRouter.POST("", handler.Create())
	productRouter.PUT("/:id", handler.Update())
	productRouter.DELETE("/:id", handler.Delete())
}
