package templates

import (
	"html/template"
	"github.com/ZacharyJacobCollins/GroupOrganization/models"
	"net/http"
)



func RenderPage(w http.ResponseWriter, templ string, p models.Page) {
	t := template.Must(template.ParseFiles("templates/"+ templ +".html"))
	t.Execute(w, p)
}

func RenderAll(w http.ResponseWriter, p models.Page) {
	all := template.Must(template.ParseGlob("/templates"))
	all.Execute(w, p)
}