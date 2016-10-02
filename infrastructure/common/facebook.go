package common

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
	dtos "userservices/DTOs"

	"golang.org/x/oauth2"
)

type FBScope struct {
	Email        string `json:"email"`
	UserBirthday string `json:"birthday"`
	Name         string `json:"name"`
	UserId       string `json:"id"`
}

var FBConfigs = oauth2.Config{
	ClientID:     FB_CLIENT_ID,
	ClientSecret: FB_CLIENT_SECRECT,
	Scopes:       []string{"email", "user_birthday", "user_location", "user_about_me"},
	Endpoint: oauth2.Endpoint{
		AuthURL:  FB_AUTH_URL,
		TokenURL: FB_TOKEN_URL,
	},
	RedirectURL: FB_REDIRECT_URL,
}

func FBGraphURL(accessToken string) string {
	fmt.Printf("%s", FB_FETCH_INFO_URL)
	return FB_FETCH_INFO_URL + "?fields=" + FB_FIELDS + "&access_token=" + url.QueryEscape(accessToken)
}

func GetFBUserInfo(accessToken string) dtos.UserDTO {
	fbResp := FBScope{}
	var fbGraphUrl = FBGraphURL(accessToken)

	// Request to fb server.
	rs, _ := http.Get(fbGraphUrl)
	defer rs.Body.Close()
	if err := json.NewDecoder(rs.Body).Decode(&fbResp); err != nil {
		fmt.Println(err)
	}

	return MapFBScopeToDTO(fbResp)
}

func MapFBScopeToDTO(fbscope FBScope) dtos.UserDTO {
	dob, err := time.Parse("01/02/2006", fbscope.UserBirthday)
	Check(err)

	return dtos.UserDTO{
		FullName: fbscope.Name,
		FBId:     fbscope.UserId,
		Dob:      dob,
	}
}
