package services

import (

	"bookApp/config"
	"bookApp/dto"
	"bookApp/errs"
	"bookApp/repository"
	"bookApp/utils"
)

type AuthorServices interface {

	Login(*dto.Author) (*dto.Authentication, *errs.AppError)
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

func (a DefaultAuthorServices) Login(author *dto.Author) (*dto.Authentication, *errs.AppError) {

	if len(author.Email) == 0 && len(author.Phone) == 0 {

		return nil, errs.BadRequestError("Lost information!!!")
	}
	if len(author.Email) > 0 && len(author.Phone) > 0 {

		return nil, errs.BadRequestError("Only send email or phone!!!")
	}
	if len(author.Email) > 0 {

		err := dto.ValidateEmail(author.Email)
		if err != nil {

			return nil, err
		}
	}
	err := dto.ValidatePhone(author.Phone)
	if err != nil {

		return nil, err
	}
	
	res, err := a.repo.Login(author.Email,author.Phone)
	if err != nil {

		return nil, err
	}
	if !utils.CheckPasswordHash(res.Password, author.Password) {

		return nil, errs.BadRequestError("Password is incorrect!!!")
	}
	myMap := utils.StructToMap(dto.AuthorModelToAuthorDto(res))
	if myMap == nil {

		return nil, errs.ErrorData()
	}

	accessToken, err := config.NewAccessJsonWebToken(myMap)
	if err != nil {

		return nil, err
	}

	refreshToken, err := config.NewRefreshJsonWebToken(myMap)
	if err != nil {

		return nil, err
	}

	return &dto.Authentication{

		AccessToken: *accessToken,
		RefreshToken: *refreshToken,
		Message: "Login Success!!!",
	}, nil	
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

	err := dto.ValidateEmail(author.Email)
	if err != nil {

		return err
	}
	err = dto.ValidatePhone(author.Phone)
	if err != nil {

		return err
	}
	err = dto.CheckPassWord(author.Password)
	if err != nil {

		return err
	}
	hassPash, err := utils.HashPassword(author.Password)
	if err != nil {

		return err
	}
	author.Password = hassPash
	err = dto.CheckName(author.Name)
	if err != nil {

		return err
	}
	res, err := a.repo.GetByEmail(author.Email)
	if err != nil {

		return err
	}
	if res != nil {

		return errs.BadRequestError("Email already exists!!!")
	}
	res, err = a.repo.GetByPhone(author.Phone)
	if err != nil {

		return err
	}
	if res != nil {

		return errs.BadRequestError("Phone number already exists!!!")
	}
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
	err = dto.ValidateEmail(author.Email)
	if err != nil {

		return err
	}
	err = dto.ValidatePhone(author.Phone)
	if err != nil {

		return err
	}
	err = dto.CheckName(author.Name)
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
	
	return res.Authors, nil
}