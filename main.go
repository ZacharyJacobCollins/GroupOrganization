package main

import (
	"net/http"
	"regexp"

	"github.com/zacharyjacobcollins/GroupOrganization/handlers"
	"github.com/zacharyjacobcollins/GroupOrganization/callback"
	"github.com/gorilla/mux"
)
var validPath = regexp.MustCompile("^/(login|chat)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/callback", callback.CallbackHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.Handle("/",  http.FileServer(http.Dir("./html")))
	r.HandleFunc("/login", handlers.LoginHandler)
	r.HandleFunc("/test",  handlers.TestHandler)
	http.ListenAndServe(":1337", nil)
}