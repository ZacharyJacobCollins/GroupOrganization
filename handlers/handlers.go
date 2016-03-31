package handlers

import (
	"net/http"
	"html/template"
	"github.com/zacharyjacobcollins/GroupOrganization/models"
)

//inject a user template
func renderTemplate(w http.ResponseWriter, tmpl string, u models.User ) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, u)
}



