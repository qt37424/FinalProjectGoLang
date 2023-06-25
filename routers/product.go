package routers

import (
	handlers "FinalProject/handlers/products"
	"FinalProject/middlewares"

	"github.com/gin-gonic/gin"
)

func (r *Router) AddProductRouter(apiRouter *gin.RouterGroup) {
	productRouter := apiRouter.Group("products")
	productRouter.Use(middlewares.CheckjwtToken()) // middleware checking Login
	handler := handlers.ProductHandler{Db: r.DB}

	productRouter.GET("", middlewares.CheckIsAdmin(), handler.GetList())
	productRouter.GET("/:id", middlewares.CheckIsAdmin(), handler.GetDetail())
	productRouter.POST("", middlewares.CheckIsAdmin(), handler.Create())
	productRouter.PUT("/:id", middlewares.CheckIsAdmin(), handler.Update())
	productRouter.DELETE("/:id", middlewares.CheckIsAdmin(), handler.Delete())
}
