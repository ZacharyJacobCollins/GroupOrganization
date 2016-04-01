package main

import (
	"github.com/ZacharyJacobCollins/GroupOrganization/models"
)

//Global Variable for organization regarding
var Organization = models.Organization{"Eastern Acm", ""}

func main() {
	//Initialize Server
	Serve()
}