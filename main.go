package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/ZacharyJacobCollins/GroupOrganization/handlers"
)

func main() {
	var r = mux.NewRouter()
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.Handle("/",  http.FileServer(http.Dir("./html")))
	
	r.HandleFunc("/internal", handlers.InternalPageHandler)
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/logout", handlers.LogoutHandler).Methods("POST")
	http.ListenAndServe(":1337", r)
}