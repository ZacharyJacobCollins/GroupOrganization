//TODO User needs to be generated and fields filled in on login.
//TODO on all renderings make sure name is capitalized.
//TODO modularize everything

package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/ZacharyJacobCollins/GroupOrganization/handlers"
	"github.com/ZacharyJacobCollins/GroupOrganization/login"
)

//Initialize Router
func init() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	r.HandleFunc("/home", handlers.IndexPageHandler)
	r.Handle("/",  http.FileServer(http.Dir("./html")))
	http.Handle("/", r)
	http.ListenAndServe(":1337", nil)
}

//Global Variable for organization is set in structs
func main() {

}