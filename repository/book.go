package repository

import (

	"bookApp/errs"
	"bookApp/models"
	"database/sql"
	"fmt"
)

type BookRepository interface {

	List() ([]*models.Book, *errs.AppError)
	Create(*models.Book) (*errs.AppError)
	Delete(int) (*errs.AppError)
	Update(*models.Book) (*errs.AppError)
	GetById(int) (*models.Book, *errs.AppError)
	Filter(string, int, int) ([]*models.Book, *errs.AppError)
	AddCategory(int, int) (*errs.AppError)
	DeleteCategory(int, int) (*errs.AppError)
	AddAuthor(int, int) (*errs.AppError)
	DeleteAuthor(int, int) (*errs.AppError)
}

type DefaultBookRepository struct {

	db *sql.DB
}

func NewBookRepository(db *sql.DB) BookRepository {

	return DefaultBookRepository{
		db: db,
	}
}


func (b DefaultBookRepository) List() ([]*models.Book, *errs.AppError) {

	res, err := b.db.Query("SELECT * FROM book")

	if err != nil {

		return nil, errs.ErrorGetData()
	}

	var books []*models.Book
	for res.Next() {

		var bookDb = new(models.BookDb)
		err := res.Scan(&bookDb.Book_id, &bookDb.Name, &bookDb.Description)
		if err != nil {

			return nil, errs.ErrorReadData()
		}
		authorRepository := NewAuthorRepository(b.db)
		authors, e := authorRepository.GetByBookId(bookDb.Book_id)
		if e != nil {

			return nil, e
		}

		categoryRepository := NewCategoryReponsitory(b.db)
		categories, e := categoryRepository.GetByBookId(bookDb.Book_id)
		if e != nil {

			return nil, e
		}
		book := &models.Book{

			Book_id: bookDb.Book_id,
			Name: bookDb.Name,
			Description: bookDb.Description,
			Authors: authors,
			Categories: categories,
		}

		books = append(books, book)
	}

	return books, nil
}

func (b DefaultBookRepository) Create(newbook *models.Book) (*errs.AppError) {

	query := fmt.Sprintf("INSERT INTO book(name, description) VALUES('%s', '%s')", newbook.Name, newbook.Description)
	_, err := b.db.Query(query)

	if err != nil {

		return errs.ErrorInsertData()
	}

	return nil
}

func (b DefaultBookRepository) Update(book *models.Book) (*errs.AppError) {

	query := fmt.Sprintf("UPDATE book SET name = '%s', description = '%s' WHERE book_id = %d", book.Name, book.Description, book.Book_id)
	_, err := b.db.Query(query)

	if err != nil {

		return errs.ErrorUpdateData()
	}

	return nil
}

func (b DefaultBookRepository) Delete(book_id int) *errs.AppError {

	query := fmt.Sprintf("DELETE FROM book_author WHERE book_id = %d", book_id)
	_, err := b.db.Query(query)

	if err != nil {

		return errs.ErrorDeleteData()
	}

	query = fmt.Sprintf("DELETE FROM book_category WHERE book_id = %d", book_id)
	_, err = b.db.Query(query)

	if err != nil {

		return errs.ErrorDeleteData()
	}

	query = fmt.Sprintf("DELETE FROM book WHERE book_id = %d", book_id)
	_, err = b.db.Query(query)

	if err != nil {

		return errs.ErrorDeleteData()
	}

	return nil
}

func (b DefaultBookRepository) GetById(book_id int) (*models.Book, *errs.AppError) {

	query := fmt.Sprintf("SELECT * FROM book WHERE book_id = %d", book_id)
	res, err := b.db.Query(query)
	if err != nil {

		return nil, errs.ErrorGetData()
	}

	var book *models.Book
	for res.Next() {

		var bookDb = new(models.BookDb)
		err := res.Scan(&bookDb.Book_id, &bookDb.Name, &bookDb.Description)
		if err != nil {

			return nil, errs.ErrorReadData()
		}
		authorRepository := NewAuthorRepository(b.db)
		authors, e := authorRepository.GetByBookId(bookDb.Book_id)
		if e != nil {

			return nil, e
		}

		categoryRepository := NewCategoryReponsitory(b.db)
		categories, e := categoryRepository.GetByBookId(bookDb.Book_id)
		if e != nil {

			return nil, e
		}
		book = &models.Book{

			Book_id: bookDb.Book_id,
			Name: bookDb.Name,
			Description: bookDb.Description,
			Authors: authors,
			Categories: categories,
		}
	}

	return book, nil
}

func (b DefaultBookRepository) Filter(search string, author_id int, category_id int) ([]*models.Book, *errs.AppError) {

	queryAuthor := ""
	if author_id != 0 {

		queryAuthor = fmt.Sprintf("JOIN book_author ba ON ba.book_id = b.book_id AND ba.author_id = %d", author_id)
	}
	queryCategory := ""
	if category_id != 0 {

		queryCategory = fmt.Sprintf("JOIN book_category bc ON bc.book_id = b.book_id AND bc.category_id = %d", category_id)
	}
	querySearch := "WHERE b.name LIKE '" + search + "%'"
	query := fmt.Sprintf("SELECT b.* FROM book b %s %s %s", queryAuthor, queryCategory, querySearch)
	//fmt.Println(query)
	res, e := b.db.Query(query)
	if e != nil {

		return nil, errs.ErrorGetData()
	}
	var books []*models.Book
	for res.Next() {

		var bookDb = new(models.BookDb)
		err := res.Scan(&bookDb.Book_id, &bookDb.Name, &bookDb.Description)
		if err != nil {

			return nil, errs.ErrorReadData()
		}
		authorRepository := NewAuthorRepository(b.db)
		authors, e := authorRepository.GetByBookId(bookDb.Book_id)
		if e != nil {

			return nil, e
		}

		categoryRepository := NewCategoryReponsitory(b.db)
		categories, e := categoryRepository.GetByBookId(bookDb.Book_id)
		if e != nil {

			return nil, e
		}
		book := &models.Book{

			Book_id: bookDb.Book_id,
			Name: bookDb.Name,
			Description: bookDb.Description,
			Authors: authors,
			Categories: categories,
		}

		books = append(books, book)
	}

	return books, nil
}

func (b DefaultBookRepository) AddCategory(book_id int, category_id int) (*errs.AppError) {

	query := fmt.Sprintf("SELECT * FROM book_category WHERE book_id = %d AND category_id = %d", book_id, category_id)
	res, e := b.db.Query(query)
	if e != nil {

		return errs.InternalServerError("Server Error")
	}
	cnt := 0
	for res.Next() {

		cnt += 1
	}
	if cnt > 0 {

		return errs.BadRequestError("Data already exist")
	}
	query = fmt.Sprintf("INSERT INTO book_category(book_id, category_id) VALUES(%d, %d)", book_id, category_id)
	_, e = b.db.Query(query)
	if e != nil {

		return errs.BadRequestError("Have no category or book")
	}
	return nil
}

func (b DefaultBookRepository) AddAuthor(book_id int, author_id int) (*errs.AppError) {

	query := fmt.Sprintf("SELECT * FROM book_author WHERE book_id = %d AND author_id = %d", book_id, author_id)
	res, e := b.db.Query(query)
	if e != nil {

		return errs.InternalServerError("Server Error")
	}
	cnt := 0
	for res.Next() {

		cnt += 1
	}
	if cnt > 0 {

		return errs.BadRequestError("Data already exist")
	}
	query = fmt.Sprintf("INSERT INTO book_author(book_id, author_id) VALUES(%d, %d)", book_id, author_id)
	_, e = b.db.Query(query)
	if e != nil {

		return errs.BadRequestError("Have no author or book")
	}
	return nil
}

func (b DefaultBookRepository) DeleteAuthor(book_id int, author_id int) (*errs.AppError) {

	query := fmt.Sprintf("SELECT * FROM book_author WHERE book_id = %d AND author_id = %d", book_id, author_id)
	res, e := b.db.Query(query)
	if e != nil {

		return errs.InternalServerError("Server Error")
	}
	cnt := 0
	for res.Next() {

		cnt += 1
	}
	if cnt == 0 {

		return errs.ErrorDataNotSurvive()
	}
	query = fmt.Sprintf("DELETE FROM book_author WHERE book_id = %d AND author_id = %d", book_id, author_id)
	_, e = b.db.Query(query)
	if e != nil {

		return errs.InternalServerError("Server Error")
	}
	return nil
}

func (b DefaultBookRepository) DeleteCategory(book_id int, category_id int) (*errs.AppError) {

	query := fmt.Sprintf("SELECT * FROM book_category WHERE book_id = %d AND category_id = %d", book_id, category_id)
	res, e := b.db.Query(query)
	if e != nil {

		return errs.InternalServerError("Server Error")
	}
	cnt := 0
	for res.Next() {

		cnt += 1
	}
	if cnt == 0 {

		return errs.ErrorDataNotSurvive()
	}
	query = fmt.Sprintf("DELETE FROM book_category WHERE book_id = %d AND category_id = %d", book_id, category_id)
	_, e = b.db.Query(query)
	if e != nil {

		return errs.InternalServerError("Server Error")
	}
	return nil
}