package login

import (
	"net/http"
	"github.com/gorilla/securecookie"
	"github.com/ZacharyJacobCollins/GroupOrganization/templates"
	"github.com/ZacharyJacobCollins/GroupOrganization/models"
)

var cookieGenerator = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	templates.RenderPage(w, "login")
}

//Index.html is the login for now.
func AuthHandler(w http.ResponseWriter, r *http.Request) {
	username 	:= r.FormValue("name")
	password 	:= r.FormValue("password")
	picture := r.FormValue("picture")
	u := models.CreateUser(username, "", picture, password)
	redirectTarget := "/"

	//TODO change here to authentication method.
	if username != "" && password != "" {
		// .. check credentials ..
		setSession(u, w)
		//RENDER ALL PAGES on successful authentication.
		templates.RenderAll(w, u)
		redirectTarget = "/home.html"
	}
	http.Redirect(w, r, redirectTarget, 302)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	clearSession(w)
	http.Redirect(w, r, "/", 302)
}
