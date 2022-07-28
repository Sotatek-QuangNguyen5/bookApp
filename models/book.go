package models

type Book struct {

	Book_id int 				`json:"book_id"`
	Name string 				`json:"name"`
	Description string 			`json:"description"`
	Authors []*Author			`json:"authors"`
	Categories []*Category		`json:"categories"`
}

type BookDb struct {

	Name string 				`json:"name"`
	Description string 			`json:"description"`
	Book_id int 				`json:"book_id"`
}
