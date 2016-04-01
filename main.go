//TODO User needs to be generated and fields filled in on login.
//TODO on all renderings make sure name is capitalized.
//TODO modularize everything, change directories of templates.

package main

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
	r.HandleFunc("/",  login.LoginHandler).Methods("GET")
	r.HandleFunc("/login", login.AuthHandler).Methods("POST")
	r.HandleFunc("/logout", login.LogoutHandler).Methods("POST")

	//Facebook Login routes
	r.HandleFunc("/FBLogin", login.FBLogin)


	http.Handle("/", r)
	http.ListenAndServe(":1337", nil)
}

//Global Variable for organization is set in structs
func main() {

}