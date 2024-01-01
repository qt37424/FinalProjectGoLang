package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Router struct {
	Server *gin.Engine
	DB     *gorm.DB
}

func (r *Router) Init() {
	apiRouter := r.Server.Group("api")
	r.AddUserRouter(apiRouter)
	r.AddProductRouter(apiRouter)
}
