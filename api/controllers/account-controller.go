package controllers

import (
	"fmt"
	"net/http"
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
	authCodeURL := common.FBConfigs.AuthCodeURL("foo")
	fmt.Println(authCodeURL)

	c.HTML(http.StatusOK, "loginviafb.html", gin.H{
		"authCodeURL": authCodeURL,
	})
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
	fbData.FBAccessToken = token.AccessToken

	matchedUser := userRepository.GetOneBy("fbId", fbData.FBId)
	if matchedUser.Id != "" {
		matchedUser.MapFromFBUser(fbData)
		userRepository.UpdateById(matchedUser.Id, matchedUser)
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": common.LOGIN_WITH_FB_SUCCESS,
		})
	}

	if matchedUser.Id == "" {
		rs := userRepository.CreateUserWithFB(fbData)
		respMessage := common.LOGIN_WITH_FB_SUCCESS
		if !rs {
			respMessage = common.INTERNAL_EXCEPTION_MESSAGE
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"data": respMessage,
		})
	}
}

func getFBUser(c *gin.Context) {
	fbId := c.Query("fbId")
	matchedUser := userRepository.GetOneBy("fbId", fbId)
	data := common.FetchInfo(matchedUser.FBAccessToken)
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
