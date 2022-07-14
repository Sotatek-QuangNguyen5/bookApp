package services

import (
	
	"bookApp/dto"
	"bookApp/errs"
	"bookApp/repository"
)

type BookServices interface {

	ListBook() ([]*dto.Book, *errs.AppError)
	DeleteBook(*dto.Book) (*errs.AppError)
	CreateBook(*dto.Book) (*errs.AppError)
	UpdateBook(*dto.Book) (*errs.AppError)
	GetByIdBook(*dto.Book) (*dto.Book, *errs.AppError)
	FilterBook(*dto.FilterBook) ([]*dto.Book, *errs.AppError)
	AddCategory(*dto.BookCategory) (*errs.AppError)
	AddAuthor(*dto.BookAuthor) (*errs.AppError)
	DeleteCategory(*dto.BookCategory) (*errs.AppError)
	DeleteAuthor(*dto.BookAuthor) (*errs.AppError)
}

type DefaultBookServices struct {

	repo repository.BookRepository
}


func NewBookServices(repo repository.BookRepository) BookServices {

	return DefaultBookServices{

		repo: repo,
	}
}

func (b DefaultBookServices) ListBook() ([]*dto.Book, *errs.AppError) {

	books, err := b.repo.List()
	if err != nil {

		return nil, err
	}

	dtoBooks := dto.BooksModelToBooksDto(books)
	return dtoBooks, nil
}

func (b DefaultBookServices) DeleteBook(book *dto.Book) (*errs.AppError) {

	err := dto.CheckID(book.Book_id)
	if err != nil {

		return err
	}
	return b.repo.Delete(book.Book_id)
}

func (b DefaultBookServices) CreateBook(book *dto.Book) (*errs.AppError) {
	
	return b.repo.Create(dto.BookDtoToBookModel(book))
}

func (b DefaultBookServices) UpdateBook(book *dto.Book) (*errs.AppError) {

	err := dto.CheckID(book.Book_id)
	if err != nil {

		return err
	}
	return b.repo.Update(dto.BookDtoToBookModel(book))
}

func (b DefaultBookServices) GetByIdBook(book *dto.Book) (*dto.Book, *errs.AppError) {

	e := dto.CheckID(book.Book_id)
	if e != nil {

		return nil, e
	}
	res, err := b.repo.GetById(book.Book_id)
	if err != nil {

		return nil, err
	}
	bookDto := dto.BookModelToBookDto(res)
	return bookDto, nil
}

func (b DefaultBookServices) FilterBook(filter *dto.FilterBook) ([]*dto.Book, *errs.AppError) {

	res, err := b.repo.Filter(filter.Search, filter.Author_id, filter.Category_id)
	if err != nil {

		return nil, err
	}
	// if len(res) == 0 {

	// 	return nil, errs.ErrorDataNotSurvive()
	// }
	return dto.BooksModelToBooksDto(res), nil
}

func (b DefaultBookServices) AddCategory(bc *dto.BookCategory) (*errs.AppError) {

	return b.repo.AddCategory(bc.Book_id, bc.Category_id)
}

func (b DefaultBookServices) AddAuthor(ba *dto.BookAuthor) (*errs.AppError) {

	return b.repo.AddAuthor(ba.Book_id, ba.Author_id)
}

func (b DefaultBookServices) DeleteCategory(bc *dto.BookCategory) (*errs.AppError) {

	return b.repo.DeleteCategory(bc.Book_id, bc.Category_id)
}

func (b DefaultBookServices) DeleteAuthor(ba *dto.BookAuthor) (*errs.AppError) {

	return b.repo.DeleteAuthor(ba.Book_id, ba.Author_id)
}