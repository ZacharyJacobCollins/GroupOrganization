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

func RenderCSS(w http.ResponseWriter, templ string, o models.Organization) {
	//TODO for range of .css files need to parse.
}

//func RenderAll(w http.ResponseWriter, templ string, p models.Page) {
//
//}