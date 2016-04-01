package main

import (
	"github.com/gorilla/mux"
	"net/http"

	"github.com/ZacharyJacobCollins/GroupOrganization/handlers"

)

func Serve() {
	var r = mux.NewRouter()
	//Serve front-end assets.
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	r.Handle("/",  http.FileServer(http.Dir("./html")))
	r.HandleFunc("/internal", handlers.IndexPageHandler)
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/logout", handlers.LogoutHandler).Methods("POST")
	http.Handle("/", r)
	http.ListenAndServe(":1337", nil)
}