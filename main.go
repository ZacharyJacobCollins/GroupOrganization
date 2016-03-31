package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var r = mux.NewRouter()
	r.HandleFunc("/", indexPageHandler)
	r.HandleFunc("/internal", internalPageHandler)

	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/logout", logoutHandler).Methods("POST")

	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}