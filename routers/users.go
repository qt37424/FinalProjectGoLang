package routers

import (
	handlers "FinalProject/handlers/users"
	"FinalProject/middlewares"

	"github.com/gin-gonic/gin"
)

func (r *Router) AddUserRouter(apiRouter *gin.RouterGroup) {
	userRouter := apiRouter.Group("users")
	handler := handlers.UserHandler{Db: r.DB}

	userRouter.POST("/sign-up", handler.Register())
	userRouter.POST("/login", middlewares.IsLogined(), handler.Login())
	userRouter.GET("/", handler.GetInfo())
}
