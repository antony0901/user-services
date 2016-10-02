package main

import (
	"net/http"
	"userservices/api/controllers"

	gin "github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// Set router as a default one provided by Gin
	router = gin.Default()
	router.LoadHTMLGlob("templates/*")
	// Initalize routes
	initializeRoute()

	// Start serving the application
	router.Run(":8080")
}

func initializeRoute() {
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Home",
		})
	})

	controllers.InitUserRoutes(router)
	controllers.InitAccountRoute(router)
}
