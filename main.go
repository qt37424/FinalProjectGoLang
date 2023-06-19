package main

import (
	router "FinalProject/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize
	loadEnv()
	server := gin.Default()

	db := connectDb()
	router := router.Router{Server: server, DB: db}
	router.Init()

	server.Run(":3000") // listen and serve on 0.0.0.0:3000
}
