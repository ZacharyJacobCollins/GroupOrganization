package login

import (
	"fmt"
	"github.com/antonholmquist/jason"
	"net/http"
	"strconv"
	"strings"
	"github.com/ZacharyJacobCollins/GroupOrganization/models"
	"log"
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
	fmt.Println("GetAccessToken")
	//https://graph.facebook.com/oauth/access_token?client_id=YOUR_APP_ID&redirect_uri=YOUR_REDIRECT_URI&client_secret=YOUR_APP_SECRET&code=CODE_GENERATED_BY_FACEBOOK
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

func FBLogin(w http.ResponseWriter, r *http.Request) {
	// grab the code fragment
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	code := r.FormValue("code")
	ClientId := "" // change this to yours
	ClientSecret := ""
	RedirectURL := "http://<domain name and don't forget port number if you use one>/FBLogin"
	accessToken := GetAccessToken(ClientId, code, ClientSecret, RedirectURL)
	response, err := http.Get("https://graph.facebook.com/me?access_token=" + accessToken.Token)
	// handle err. You need to change this into something more robust
	// such as redirect back to home page with error message
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	str := readHttpBody(response)
	// dump out all the data
	// w.Write([]byte(str))
	user, _ := jason.NewObjectFromBytes([]byte(str))
	id, _ := user.GetString("id")
	email, _ := user.GetString("email")
	bday, _ := user.GetString("birthday")
	fbusername, _ := user.GetString("username")
	w.Write([]byte(fmt.Sprintf("Username %s ID is %s and birthday is %s and email is %s<br>", fbusername, id, bday, email)))
	img := "https://graph.facebook.com/" + id + "/picture?width=180&height=180"
	w.Write([]byte("Photo is located at " + img + "<br>"))
	// see https://www.socketloop.com/tutorials/golang-download-file-example on how to save FB file to disk
	w.Write([]byte("<img src='" + img + "'>"))
}
