package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/zacharyjacobcollins/GroupOrganization/handlers"
)

func main() {
	var r = mux.NewRouter()
	r.HandleFunc("/", handlers.IndexPageHandler)
	r.HandleFunc("/internal", InternalPageHandler)

	r.HandleFunc("/login", LoginHandler).Methods("POST")
	r.HandleFunc("/logout", LogoutHandler).Methods("POST")

	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}