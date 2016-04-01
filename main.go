//TODO User needs to be generated and fields filled in on login.
//TODO on all renderings make sure name is capitalized.
//TODO modularize everything, change directories of templates.
//TODO randomize non-pic selection
//TODO sanitize all inputs

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
	r.HandleFunc("/logout", login.LogoutHandler)

	//Facebook Login routes
	r.HandleFunc("/FBLogin", login.FBLoginHandler)


	http.Handle("/", r)
	http.ListenAndServe(":1337", nil)
}

//Global Variable for organization is set in Structs
func main() {

}