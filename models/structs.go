package structs

type User struct {
	Name string
	Password string
	BlogPosts []string
	Messages []string
}

type Organization struct {
	Name		string
	HexColor 	string
}