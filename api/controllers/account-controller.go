package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	dtos "userservices/DTOs"
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

func loginViaFB(c *gin.Context) {
	authCodeURL := common.FBConfigs.AuthCodeURL("")
	rs, err := http.Get(authCodeURL)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rs)
}

func auth(c *gin.Context) {
	code := c.Query("code")
	token, err := common.FBConfigs.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, common.BAD_REQUEST_MESSAGE)
	}

	dto := dtos.UserDTO{
		FBAccessToken: token.AccessToken,
	}
	user := models.User{}
	models.MapToUser(dto, &user)

}

func getFBUser(c *gin.Context) {
	fbResp := common.FBScope{}
	accessToken := c.Query("accessToken")
	var fbGraphUrl = common.FBGraphURL(accessToken)
	fmt.Println(fbGraphUrl)
	rs, _ := http.Get(fbGraphUrl)
	defer rs.Body.Close()
	if err := json.NewDecoder(rs.Body).Decode(&fbResp); err != nil {
		fmt.Println(err)
	}
	data := common.MapFBScopeToDTO(fbResp)
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
