package templates

import (
	"html/template"
	"net/http"
	"github.com/ZacharyJacobCollins/GroupOrganization/models"
)

func RenderPage(w http.ResponseWriter, templ string) {
	t := template.Must(template.ParseFiles("templates/"+templ+".html"))
	t.Execute(w, models.GlobalOrganization)
}

//Renders all templates server side.  Called after auth.
func RenderAll(w http.ResponseWriter, u models.User) {
	p := models.Page{u, models.GlobalOrganization}
	t := template.Must(template.ParseFiles(
		"templates/home.html",
	))
	t.Execute(w, p)
}