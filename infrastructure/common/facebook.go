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
	Picture      struct {
		Data struct {
			URL string `json:"url"`
		} `json:"data"`
	} `json:"picture"`
	Photos struct {
		Data []struct {
			CreatedTime string `json:"created_time"`
			Name        string `json:"name"`
			Id          string `json:"id"`
		} `json:"data"`
	} `json:"photos"`
}

var FBConfigs = oauth2.Config{
	ClientID:     FB_CLIENT_ID,
	ClientSecret: FB_CLIENT_SECRECT,
	Endpoint: oauth2.Endpoint{
		AuthURL:  FB_AUTH_URL,
		TokenURL: FB_TOKEN_URL,
	},
	RedirectURL: FB_REDIRECT_URL,
	Scopes:      []string{"email", "user_birthday", "user_location", "user_about_me", "public_profile", "user_photos"},
}

func FBGraphURL(accessToken string) string {
	return FB_FETCH_INFO_URL + "?fields=" + url.QueryEscape(FB_FIELDS) + "&access_token=" + url.QueryEscape(accessToken)
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

// FetchInfo request to Facebook to get basic info
func FetchInfo(accessToken string) FBScope {
	fbResp := FBScope{}
	fbGraphUrl := FBGraphURL(accessToken)
	// fmt.Println(fbGraphUrl)
	// fbGraphUrl = "https://graph.facebook.com/v2.7/me?fields=id%2Cname%2Cbirthday%2Ccontext%2Cemail&access_token=EAACEdEose0cBAI1uZAYpyReinhhDC0JvIksFijaaryI59dBQZAxFJdRcYm5Ylv6eZBh3a7D18F61NM9TbCOKtya15O1B0ePsKZBmCX8mdihYTqQme9qCA9pRPJ4IZBYxJlkXfeX6TMLgvRnG4zfL62FZCwaYzxjKogDTaydSZANqwZDZD"
	fmt.Println(fbGraphUrl)

	// Request to fb server.
	rs, _ := http.Get(fbGraphUrl)
	defer rs.Body.Close()

	err := json.NewDecoder(rs.Body).Decode(&fbResp)
	if err != nil {
		Check(err)
	}

	return fbResp
}

func MapFBScopeToDTO(fbscope FBScope) dtos.UserDTO {
	dob, err := time.Parse("01/02/2006", fbscope.UserBirthday)
	Check(err)

	return dtos.UserDTO{
		FullName: fbscope.Name,
		FBId:     fbscope.UserId,
		Dob:      dob,
		Email:    fbscope.Email,
		Picture:  fbscope.Picture.Data.URL,
	}
}
