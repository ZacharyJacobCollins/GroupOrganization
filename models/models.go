package models

import (
	"golang.org/x/oauth2"
	"github.com/ZacharyJacobCollins/GroupOrganization/utils"
)

var GlobalOrganization = Organization{
	"Eastern Acm",
	"Acm",
	"https://netid.emich.edu/cas/favicon.ico",
	"#000046",
	GenerateFacebookLink(),
}

func GenerateFacebookLink() string {
	//w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fbConfig := &oauth2.Config{
		ClientID:     "219381195087194",
		ClientSecret: "c292c1950b78296d55d9208cd3ad1d19",
		RedirectURL:  "http://localhost:1337/FBLogin",
		Scopes:       []string{"email", "user_birthday", "user_location", "user_about_me"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.facebook.com/dialog/oauth",
			TokenURL: "https://graph.facebook.com/oauth/access_token",
		},
	}
	return fbConfig.AuthCodeURL("")
}


//Creates user and ensures profile picture.
func CreateUser(username string, position string, pictureUrl string, password string) User{
	//Make sure name is capitalized
	username = utils.Capitalize(username)

	if (username == "") {
		username = "EMU Student"
	}
	if (position == "") {
		position = "President"
	}
	if (pictureUrl == "") {
		pictureUrl = "http://orig00.deviantart.net/9b6c/f/2013/124/2/0/arch_linux_logo_by_nintenmario-d6419p8.png"
	}
	if (password == "") {
		password = "test"
	}
	return User{username, position, pictureUrl, password}
}

type Page struct {
	U		User
	O		Organization
}

type User struct {
	Name 			string
	Position 		string
	Picture			string
	Password 		string
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