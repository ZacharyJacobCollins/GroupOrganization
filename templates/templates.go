package templates

import (
	"html/template"
	"net/http"
	"github.com/ZacharyJacobCollins/GroupOrganization/models"
)

//TODO needs to be global in core.
var GlobalOrganization = models.Organization{
	"Eastern Acm",
	"Acm",
	"https://netid.emich.edu/cas/favicon.ico",
	"#000046",
}

//username being passed in via LoginHandler.  Renders all templates server side.
func RenderAll(w http.ResponseWriter, username string) {
	u := models.User{username, "President", ""}
	p := models.Page{u, GlobalOrganization}
	t := template.Must(template.ParseFiles("templates/home.html", "templates/chat_view.html"))
	t.Execute(w, p)
}