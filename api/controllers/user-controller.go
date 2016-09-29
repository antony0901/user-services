package controllers

import (
	"net/http"
	dtos "userservices/DTOs"
	"userservices/infrastructure/common"
	"userservices/infrastructure/repositories"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes(router *gin.Engine) {
	router.GET("/user/:id", getById)
	router.GET("/getUserByName", getByName)
	router.POST("/createUser", createUser)
}

func getById(c *gin.Context) {
	userRepository := repositories.NewUserRepository()

	id := c.Param("id")
	matchedUser := userRepository.GetUserById(id)

	c.JSON(http.StatusOK, gin.H{
		"data": matchedUser,
	})
}

func getByName(c *gin.Context) {
	userRepository := repositories.NewUserRepository()

	name := c.Query("name")

	matchedUser := userRepository.GetUserByName(name)

	c.JSON(http.StatusOK, gin.H{
		"data": matchedUser,
	})
}

func createUser(c *gin.Context) {
	// Bind request body to User object.
	newUser := dtos.UserDTO{}
	err := c.Bind(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.BAD_REQUEST_MESSAGE)
	}

	userRepository := repositories.NewUserRepository()

	userRepository.CreateUser(newUser)
	c.JSON(http.StatusOK, gin.H{
		"data": "create user successfully!",
	})
}
