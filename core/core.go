//TODO need to modularize routes.

package core

import (
	"net/http"
	"github.com/ZacharyJacobCollins/GroupOrganization/handlers"
	"github.com/gorilla/mux"
	"github.com/ZacharyJacobCollins/GroupOrganization/login"
)


func init() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	r := mux.NewRouter()



	r.HandleFunc("/home", handlers.IndexPageHandler)
	r.Handle("/",  http.FileServer(http.Dir("./html")))
	r.HandleFunc("/login", login.LoginHandler).Methods("POST")
	r.HandleFunc("/logout", login.LogoutHandler).Methods("POST")



	http.Handle("/", r)
	http.ListenAndServe(":1337", nil)
}