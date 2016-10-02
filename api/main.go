package main

import (
	"userservices/api/controllers"

	gin "github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// Set router as a default one provided by Gin
	router = gin.Default()

	// Initalize routes
	initializeRoute()

	// Start serving the application
	router.Run(":8080")
}

func initializeRoute() {
	controllers.InitUserRoutes(router)
	controllers.InitAccountRoute(router)
}
