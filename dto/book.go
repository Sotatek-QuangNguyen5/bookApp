package dto

import "bookApp/models"

type Book struct {
	Book_id     int         `json:"book_id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Authors     []*Author   `json:"authors"`
	Categories  []*Category `json:"categories"`
}

type BookCategory struct {

	Book_id		int 		`json:"book_id"`
	Category_id int			`json:"category_id"`
}

type BookAuthor struct {

	Book_id		int 		`json:"book_id"`
	Author_id int			`json:"author_id"`
}

func BookModelToBookDto(bookModel *models.Book) *Book {

	if bookModel == nil {

		return nil
	}
	return &Book{

		Book_id: bookModel.Book_id,
		Name: bookModel.Name,
		Description: bookModel.Description,
		Authors: AuthorsModelToAuthorsDto(bookModel.Authors),
		Categories: CategoriesModelToCategoriesDto(bookModel.Categories),
	}
}

func BooksModelToBooksDto(booksModel []*models.Book) []*Book {

	var books []*Book
	for _, book := range booksModel {

		books = append(books, BookModelToBookDto(book))
	}
	return books
}

func BookDtoToBookModel(bookDto *Book) *models.Book {

	if bookDto == nil {

		return nil
	}
	return &models.Book{

		Book_id: bookDto.Book_id,
		Name: bookDto.Name,
		Description: bookDto.Description,
		Authors: AuthorsDtoToAuthorsModel(bookDto.Authors),
		Categories: CategoriesDtoToCategoriesModel(bookDto.Categories),
	}
}

func BooksDtoToBooksModel(booksDto []*Book) []*models.Book {

	var books []*models.Book
	for _, book := range booksDto {

		books = append(books, BookDtoToBookModel(book))
	}
	return books
}