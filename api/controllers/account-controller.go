package controllers

import (
	"fmt"
	"net/http"
	"userservices/domain/models"
	"userservices/infrastructure/common"
	"userservices/infrastructure/repositories"

	"golang.org/x/oauth2"

	"github.com/gin-gonic/gin"
)

var userRepository repositories.UserRepository

func InitAccountRoute(router *gin.Engine) {
	userRepository = repositories.NewUserRepository()
	router.GET("/loginviafacebook", loginViaFB)
	router.GET("/auth", auth)
	router.GET("/getFBUser", getFBUser)
}

// This function should be implemented on client.
func loginViaFB(c *gin.Context) {
	authCodeURL := common.FBConfigs.AuthCodeURL("")
	resp, err := http.Get(authCodeURL)
	common.Check(err)

	fmt.Println(resp)
}

// Handle RedirectURL request from Facebook
func auth(c *gin.Context) {
	code := c.Query("code")
	token, err := common.FBConfigs.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, common.BAD_REQUEST_MESSAGE)
	}

	fbData := common.GetFBUserInfo(token.AccessToken)
	matchedUser := models.User{}
	userRepository.GetOneBy("accessToken", token.AccessToken).One(&matchedUser)
	if matchedUser.Id != "" {
		models.MapToUser(fbData, &matchedUser)
		userRepository.UpdateById(matchedUser.Id, matchedUser)
	}

	rs := userRepository.CreateUserWithFB(fbData)
	respMessage := common.LOGIN_WITH_FB_SUCCESS
	if !rs {
		respMessage = common.INTERNAL_EXCEPTION_MESSAGE
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"data": respMessage,
	})
}

func getFBUser(c *gin.Context) {
	accessToken := c.Query("accessToken")
	data := common.GetFBUserInfo(accessToken)
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
