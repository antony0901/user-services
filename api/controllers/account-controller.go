package controllers

import (
	"net/http"
	dtos "userservices/DTOs"
	"userservices/infrastructure/common"
	"userservices/infrastructure/repositories"

	"golang.org/x/oauth2"

	"github.com/gin-gonic/gin"
)

func InitAccountRoute(router *gin.Engine) {
	router.GET("/loginviafacebook", loginViaFB)
}

func loginViaFB(c *gin.Context) {
	code := c.Query("code")
	token, err := common.FBConfigs.Exchange(oauth2.NoContext, code)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	newUser := dtos.UserDTO{
		FBAccessToken: token.AccessToken,
	}

	userRepository := repositories.NewUserRepository()

	userRepository.CreateUser(newUser)

	c.JSON(http.StatusOK, gin.H{
		"data": "create user successfully!",
	})
}
