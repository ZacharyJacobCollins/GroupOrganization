package handlers

import (
	"net/http"
	"github.com/ZacharyJacobCollins/GroupOrganization/models"

	"github.com/ZacharyJacobCollins/GroupOrganization/templates"
)

var GlobalOrganization = models.Organization{"Eastern Acm", "Acm", "#000046"}

func IndexPageHandler(w http.ResponseWriter, request *http.Request) {
	u := models.User{getUserName(request), ""}
	p := models.Page{u, GlobalOrganization}
	templates.RenderPage(w, "home", p)
}