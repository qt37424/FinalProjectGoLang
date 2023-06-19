package router

import (
	"FinalProject/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) AddUserRouter(apiRouter *gin.RouterGroup) {
	userRouter := apiRouter.Group("users")

	userRouter.GET("/", func(ctx *gin.Context) {
		users := &[]models.User{}
		r.DB.Find(users)
		ctx.JSON(http.StatusOK, gin.H{"data": users})
	})

	userRouter.POST("/sign-up", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "signUp"})
	})

	userRouter.POST("/login", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "login"})
	})
}
