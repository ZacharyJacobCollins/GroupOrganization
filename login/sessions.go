package login

import (
	"net/http"
	"github.com/ZacharyJacobCollins/GroupOrganization/models"
)

func getUser(request *http.Request) (user models.User) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]models.User)
		if err = cookieGenerator.Decode("session", cookie.Value, &cookieValue); err == nil {
			user = cookieValue["name"]
		}
	}
	return user
}

func setSession(user models.User, response http.ResponseWriter) {
	value := map[string]models.User{
		"user": user,
	}
	if encoded, err := cookieGenerator.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}