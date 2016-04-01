//TODO User needs to be generated and fields filled in on login.

package main

import (
	"github.com/ZacharyJacobCollins/GroupOrganization/models"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/ZacharyJacobCollins/GroupOrganization/handlers"
)

//Global Variable for organization.  Never changes.
const GlobalOrganization = models.Organization{"Eastern Acm", "Acm", "#000046"}


func main() {
	var r = mux.NewRouter()
	//Serve front-end assets.
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	r.Handle("/",  http.FileServer(http.Dir("./html")))
	r.HandleFunc("/home", handlers.IndexPageHandler)
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/logout", handlers.LogoutHandler).Methods("POST")
	http.Handle("/", r)
	http.ListenAndServe(":1337", nil)
}