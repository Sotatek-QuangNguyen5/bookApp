package dto


type FilterBook struct {

	Search string 			`json:"search"`
	Author_id int			`json:"author_id"`
	Category_id int			`json:"category_id"`
}