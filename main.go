package main

import (
	"net/http"

	"github.com/zacharyjacobcollins/GroupOrganization/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	//Handle front end assets
	r.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	r.Handle("/",  http.FileServer(http.Dir("./html")))
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	http.ListenAndServe(":1337", r)
}