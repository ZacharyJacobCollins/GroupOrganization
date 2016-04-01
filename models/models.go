package models

import (
"golang.org/x/oauth2"
)

//TODO needs to be global in core.
var GlobalOrganization = Organization{
	"Eastern Acm",
	"Acm",
	"https://netid.emich.edu/cas/favicon.ico",
	"#000046",
	GenerateAuthLink(),
}

func GenerateAuthLink() string {
	//w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fbConfig := &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		RedirectURL:  "http://<domain name and don't forget port number if you use one>/FBLogin", // change this to your webserver adddress
		Scopes:       []string{"email", "user_birthday", "user_location", "user_about_me"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.facebook.com/dialog/oauth",
			TokenURL: "https://graph.facebook.com/oauth/access_token",
		},
	}
	return fbConfig.AuthCodeURL("")
}

type Page struct {
	U		User
	O		Organization
}

type User struct {
	Name 		string
	Position 	string
	Password 	string
}

type Organization struct {
	Name			string
	//Kept to 3 characters
	Initials 		string
	Favicon			string
	HexColor 		string
	FacebookLogin 		string
}

type AccessToken struct {
	Token  string
	Expiry int64
}