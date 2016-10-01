package main

import gin "github.com/gin-gonic/gin"

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
	controllers_user.InitUserRoutes(router)
	controllers_account.InitAccountRoute(router)
}
