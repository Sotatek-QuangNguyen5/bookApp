package services

import (
	"bookApp/config"
	"bookApp/dto"
	"bookApp/errs"
	"bookApp/repository"
)

type AuthorServices interface {

	ListAuthor() ([]*dto.Author, *errs.AppError)
	DeleteAuthor(*dto.Author) (*errs.AppError)
	UpdateAuthor(*dto.Author) (*errs.AppError)
	CreateAuthor(*dto.Author) (*errs.AppError)
	GetByIdAuthor(*dto.Author) (*dto.Author, *errs.AppError)
	GetByBookId(*dto.Book) ([]*dto.Author, *errs.AppError)
}

type DefaultAuthorServices struct {

	repo repository.AuthorRepository
}

func NewAuthorServices(repo repository.AuthorRepository) AuthorServices {

	return DefaultAuthorServices{

		repo: repo,
	}
}

func (a DefaultAuthorServices) ListAuthor() ([]*dto.Author, *errs.AppError) {

	res, err := a.repo.List()
	if err != nil {

		return nil, err
	}
	authors := dto.AuthorsModelToAuthorsDto(res)
	return authors, nil
}

func (a DefaultAuthorServices) CreateAuthor(author *dto.Author) (*errs.AppError) {

	return a.repo.Create(dto.AuthorDtoToAuthorModel(author))
}

func (a DefaultAuthorServices) DeleteAuthor(author *dto.Author) (*errs.AppError) {

	err := dto.CheckID(author.Author_id)
	if err != nil {

		return err
	}
	return a.repo.Delete(author.Author_id)
}

func (a DefaultAuthorServices) UpdateAuthor(author *dto.Author) (*errs.AppError) {

	err := dto.CheckID(author.Author_id)
	if err != nil {

		return err
	}
	return a.repo.Update(dto.AuthorDtoToAuthorModel(author))
}

func (a DefaultAuthorServices) GetByIdAuthor(author *dto.Author) (*dto.Author, *errs.AppError) {

	err := dto.CheckID(author.Author_id)
	if err != nil {

		return nil, err
	}
	
	res, e := a.repo.GetById(author.Author_id)
	if e != nil {

		return nil, e
	}
	authorDto := dto.AuthorModelToAuthorDto(res)
	return authorDto, nil
}

func (a DefaultAuthorServices) GetByBookId(book *dto.Book) ([]*dto.Author, *errs.AppError) {

	e := dto.CheckID(book.Book_id)
	if e != nil {

		return nil, e
	}
	b := NewBookServices(repository.NewBookRepository(config.DB))
	res, err := b.GetByIdBook(book)
	if err != nil {

		return nil, err
	}
	// if res.Authors == nil || len(res.Authors) == 0 {

	// 	return nil, errs.ErrorDataNotSurvive()
	// }
	return res.Authors, nil
}