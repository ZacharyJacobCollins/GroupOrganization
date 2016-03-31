package handlers

import (
	"net/http"
	"html/template"

	"github.com/zacharyjacobcollins/GroupOrganization/models"
)

func renderTemplate(w http.ResponseWriter, tmpl string, ) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, _)
}

func LoginHandler(w http.ResponseWriter, r *http.Request, user string) {
	renderTemplate(w, "/templates/chat_view")
}

func ChatHandler(w http.ResponseWriter, r *http.Request) {

}


