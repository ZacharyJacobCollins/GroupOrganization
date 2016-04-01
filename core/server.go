package core

import (
	"net/http"
	"github.com/ZacharyJacobCollins/GroupOrganization/handlers"
	"github.com/gorilla/mux"
)

func init() {
	var r = mux.NewRouter()
	r.HandleFunc("/home", handlers.IndexPageHandler)
	r.Handle("/",  http.FileServer(http.Dir("./html")))
	http.Handle("/", r)
	http.ListenAndServe(":1337", nil)
}