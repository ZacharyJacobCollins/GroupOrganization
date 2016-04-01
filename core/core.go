package core

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/ZacharyJacobCollins/GroupOrganization/login"
)



func init() {
	//frontend assets
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))

	r := mux.NewRouter()

	//login routes
	r.Handle("/",  http.FileServer(http.Dir("./html")))
	r.HandleFunc("/login", login.LoginHandler).Methods("POST")
	r.HandleFunc("/logout", login.LogoutHandler).Methods("POST")


	http.Handle("/", r)
	http.ListenAndServe(":1337", nil)
}