package login

import (
	"github.com/gorilla/securecookie"
	"net/http"

	"github.com/ZacharyJacobCollins/GroupOrganization/templates"
)

var cookieGenerator = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))


func LoginHandler(w http.ResponseWriter, r *http.Request) {
	templates.RenderPage(w, "login")
}

//Index.html is the login for now.
func AuthHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	pass := r.FormValue("password")
	redirectTarget := "/"

	//TODO change here to authentication method.
	if name != "" && pass != "" {
		// .. check credentials ..
		setSession(name, w)
		//TODO call here to templating engine upon successful login.  RENDER ALL PAGES.
		templates.RenderAll(w, getUserName(r))
		redirectTarget = "/home.html"
	}
	http.Redirect(w, r, redirectTarget, 302)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	clearSession(w)
	http.Redirect(w, r, "/login", 302)
}
