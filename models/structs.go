package models

type Page struct {
	U		User
	O		Organization
}

type User struct {
	Name 		string
	Position 	string
	Password 	string
}

type Organization struct {
	Name		string
	//Kept to 3 characters
	Initials 	string
	Favicon		string
	HexColor 	string
}

