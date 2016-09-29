package common

import "golang.org/x/oauth2"

const (
	FB_CLIENT_ID      = "315627565448019"
	FB_CLIENT_SECRECT = "92abb488a9349ead127b43d3fd3c3805"
	FB_FETCH_INFO_URL = "https://graph.facebook.com/me?fields=id,name,birthday&access_token="
)

type FBScope struct {
	Email        string
	UserBirthday string
	UserLocation string
	UserAboutMe  string
}

var FBConfigs = oauth2.Config{
	ClientID:     FB_CLIENT_ID,
	ClientSecret: FB_CLIENT_SECRECT,
	Scopes:       []string{"email", "user_birthday", "user_location", "user_about_me"},
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://graph.facebook.com/oauth/authorize",
		TokenURL: "https://graph.facebook.com/oauth/access_token",
	},
	RedirectURL: "http://localhost:8080/auth",
}
