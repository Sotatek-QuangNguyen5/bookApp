package repository

import (
	"bookApp/errs"
	"bookApp/models"
	"database/sql"
	"fmt"
)

type AuthorRepository interface {

	List() ([]*models.Author, *errs.AppError)
	Create(*models.Author) (*errs.AppError)
	Delete(int) (*errs.AppError)
	Update(*models.Author) (*errs.AppError)
	GetById(int) (*models.Author, *errs.AppError)
	GetByBookId(int) ([]*models.Author, *errs.AppError)
}

type DefaultAuthorRepository struct {

	db *sql.DB
}

func NewAuthorRepository(db *sql.DB) AuthorRepository {

	return DefaultAuthorRepository{

		db: db,
	}
}

func (a DefaultAuthorRepository) List() ([]*models.Author, *errs.AppError) {

	res, err := a.db.Query("SELECT * FROM author")

	if err != nil {

		return nil, errs.ErrorGetData()
	}

	var authors []*models.Author
	for res.Next() {

		var author = new(models.Author)
		err = res.Scan(&author.Author_id, &author.Email, &author.Name, &author.Phone)
		if err != nil {

			return nil, errs.ErrorReadData()
		}

		authors = append(authors, author)
	}

	return authors, nil
}

func (a DefaultAuthorRepository) Create(author *models.Author) (*errs.AppError) {

	query := fmt.Sprintf("INSERT INTO author(name, email, phone) VALUES('%s','%s','%s')", author.Name, author.Email, author.Phone)
	_, err := a.db.Query(query)

	if err != nil {

		return errs.ErrorInsertData()
	}

	return nil
}

func (a DefaultAuthorRepository) Delete(Author_id int) (*errs.AppError) {

	_, e := a.GetById(Author_id)
	if e != nil {

		return e
	}

	query := fmt.Sprintf("DELETE FROM book_author WHERE author_id = '%d'", Author_id)
	_, err := a.db.Query(query)

	if err != nil {

		return errs.ErrorDeleteData()
	}

	query = fmt.Sprintf("DELETE FROM author WHERE author_id = '%d'", Author_id)
	_, err = a.db.Query(query)

	if err != nil {

		return errs.ErrorDeleteData()
	}

	return nil
}

func (a DefaultAuthorRepository) Update(author *models.Author) (*errs.AppError) {

	_, e := a.GetById(author.Author_id)
	if e != nil {

		return e
	}
	
	query := fmt.Sprintf("UPDATE author SET name = '%s', email = '%s', phone = '%s' WHERE author_id = '%d'", author.Name, author.Email, author.Phone, author.Author_id)
	_, err := a.db.Query(query)

	if err != nil {

		return errs.ErrorUpdateData()
	}

	return nil
}

func (a DefaultAuthorRepository) GetById(Author_id int) (*models.Author, *errs.AppError) {

	query := fmt.Sprintf("SELECT * FROM author WHERE author_id = %d", Author_id)
	res, err := a.db.Query(query)
	
	if err != nil {

		return nil, errs.ErrorGetData()
	}

	var author = new(models.Author)
	for res.Next() {

		e := res.Scan(author.Author_id, author.Email, author.Name, author.Phone)
		if e != nil {

			return nil, errs.ErrorReadData()
		}
	}
	if author.Author_id == 0 {

		return nil, errs.ErrorDataNotSurvive()
	}

	return author, nil
}

func (a DefaultAuthorRepository) GetByBookId(book_id int) ([]*models.Author, *errs.AppError) {

	query := fmt.Sprintf("SELECT a.* FROM author a JOIN book_author ba ON a.author_id = ba.author_id AND ba.book_id = %d", book_id)
	res, err := a.db.Query(query)
	if err != nil {

		return nil, errs.ErrorGetData()
	}

	var authors []*models.Author
	for res.Next() {

		var author = new(models.Author)
		err := res.Scan(&author.Author_id, &author.Name, &author.Email, &author.Phone)
		if err != nil {

			return nil, errs.ErrorReadData()
		}
		authors = append(authors, author)
	}

	return authors, nil
}