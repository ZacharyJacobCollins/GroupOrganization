package models



type Page struct {
	U		User
	O		Organization
}

type User struct {
	Name string
	Password string
}

type Organization struct {
	Name		string
	//used in logo etc.
	Shortname 	string
	HexColor 	string
}