package login

import (
	"github.com/antonholmquist/jason"
	"net/http"
	"strconv"
	"strings"
	"github.com/ZacharyJacobCollins/GroupOrganization/models"
	"log"
"github.com/ZacharyJacobCollins/GroupOrganization/templates"
)

func readHttpBody(response *http.Response) string {
	var str string
	bodyBuffer := make([]byte, 5000)
	count, err := response.Body.Read(bodyBuffer)
	for ; count > 0; count, err = response.Body.Read(bodyBuffer) {
		if err != nil {
			log.Print(err)
		}
		str += string(bodyBuffer[:count])
	}
	return str
}

//Converts a code to an Auth_Token
func GetAccessToken(client_id string, code string, secret string, callbackUri string) models.AccessToken {
	response, err := http.Get("https://graph.facebook.com/oauth/access_token?client_id=" +
	client_id + "&redirect_uri=" + callbackUri +
	"&client_secret=" + secret + "&code=" + code)

	if err == nil {
		auth := readHttpBody(response)
		var token models.AccessToken
		tokenArr := strings.Split(auth, "&")
		token.Token = strings.Split(tokenArr[0], "=")[1]
		expireInt, err := strconv.Atoi(strings.Split(tokenArr[1], "=")[1])
		if err == nil {
			token.Expiry = int64(expireInt)
		}
		return token
	}
	var token models.AccessToken
	return token
}

func FBLoginHandler(w http.ResponseWriter, r *http.Request){
	// grab the code fragment
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	code := r.FormValue("code")
	ClientId := "219381195087194" // change this to yours
	ClientSecret := "c292c1950b78296d55d9208cd3ad1d19"
	RedirectURL := "http://localhost:1337/FBLogin"
	accessToken := GetAccessToken(ClientId, code, ClientSecret, RedirectURL)
	response, err := http.Get("https://graph.facebook.com/me?access_token=" + accessToken.Token)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	str := readHttpBody(response)
	user, _ := jason.NewObjectFromBytes([]byte(str))
	//email, _ := user.GetString("email")
	//bday, _ := user.GetString("birthday")
	fbusername, _ := user.GetString("username")
	id, _ := user.GetString("id")
	img := "https://graph.facebook.com/" + id + "/picture?width=90&height=90"

	//this junk returns the user and stuff.
	u := GetFacebookUser(fbusername, "", img, "")
	log.Print(u.Name+" "+u.Picture)
	templates.RenderAll(w, u)
	http.Redirect(w, r, "/templates/home.html", 302)
}

func GetFacebookUser(fbusername string, position string, img string, password string) models.User {
	return models.CreateUser(fbusername, "", img, "")
}





