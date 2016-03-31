package main

import (
	"net/http"
	"regexp"

	"github.com/zacharyjacobcollins/GroupOrganization/handlers"
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
	// w := wiki.NewWiki();  w.Run();
	http.Handle("/assets", http.StripPrefix("/", http.FileServer(http.Dir("/assets/"))))
	http.HandleFunc("/login", (makeHandler(handlers.loginHandler)))
	http.ListenAndServe(":1337", nil)
}