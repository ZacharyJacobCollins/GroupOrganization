package handlers

import (
	"net/http"
	"github.com/ZacharyJacobCollins/GroupOrganization/models"

)


func IndexPageHandler(response http.ResponseWriter, request *http.Request) {
	u := models.User{getUserName(request)}
	p := models.Page{u, GlobalOrganization}
}