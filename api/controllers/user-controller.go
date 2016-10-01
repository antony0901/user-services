package controllers

import (
	"net/http"
	dtos "userservices/DTOs"
	"userservices/infrastructure/common"
	"userservices/infrastructure/repositories"

	"github.com/gin-gonic/gin"
)

// Define the repositories used in current controller
var ucUserRepository repositories.UserRepository

// Init routes handled in this controller.
// InitUserRoutes called in main.go
func InitUserRoutes(router *gin.Engine) {
	ucUserRepository = repositories.NewUserRepository()

	router.GET("/user/:id", getById)
	router.GET("/getUserByName", getByName)
	router.POST("/createUser", createUser)
}

// getById is handler of user/:id route.
// This one is GET method with id param
func getById(c *gin.Context) {
	id := c.Param("id")
	matchedUser := ucUserRepository.GetUserById(id)

	c.JSON(http.StatusOK, gin.H{
		"data": matchedUser,
	})
}

// getByName is handler of /getUserByName
// This one is GET method with name as a query string.
func getByName(c *gin.Context) {
	name := c.Query("name")

	matchedUser := ucUserRepository.GetUserByName(name)

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

	ucUserRepository.CreateUser(newUser)
	c.JSON(http.StatusOK, gin.H{
		"data": "create user successfully!",
	})
}
