package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	dtos "userservices/DTOs"
	"userservices/infrastructure/common"
	"userservices/infrastructure/repositories"

	"golang.org/x/oauth2"

	"github.com/gin-gonic/gin"
)

func InitAccountRoute(router *gin.Engine) {
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

	newUser := dtos.UserDTO{
		FullName:      "linh",
		FBAccessToken: token.AccessToken,
	}

	userRepository := repositories.NewUserRepository()

	userRepository.CreateUser(newUser)
}

func getFBUser(c *gin.Context) {
	me := map[string]interface{}{}
	var tempAccessToken = "EAAEfD8Tre1MBAFWTBwa3Sg2zfjU6XJKhXZB8bI2scZAjUYS9RDjeOouZCdllZAN50CHLa7cXCrih2nG67MEIPj3GGLJ7AZCsENKUsq230dsJ39yToESGfXJMv1y4aRKZBRBIcrr0ZBMGPLOmV9Hvgxbo8BLzGODx6oLN452mUjTtwZDZD"
	rs, _ := http.Get(common.FB_FETCH_INFO_URL + url.QueryEscape(tempAccessToken))
	defer rs.Body.Close()
	if err := json.NewDecoder(rs.Body).Decode(&me); err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": me,
	})
}
