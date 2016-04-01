package handlers

import (
	"net/http"
	"github.com/ZacharyJacobCollins/GroupOrganization/models"
)

var GlobalOrganization = models.Organization{
	"Eastern Acm",
	"Acm",
	"https://netid.emich.edu/cas/favicon.ico",
	"#000046",
}

func IndexPageHandler(w http.ResponseWriter, request *http.Request) {
	//u := models.User{getUserName(request), "President", ""}
	//p := models.Page{u, GlobalOrganization}
	//templates.RenderPage(w, "home", p)
}

