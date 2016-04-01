package login

import (
	"github.com/HubHtml/login"
	"github.com/gorilla/mux"
)

func init() {
	var r = mux.NewRouter()
	r.HandleFunc("/login", LoginHandler).Methods("POST")
	r.HandleFunc("/logout", LogoutHandler).Methods("POST")
}