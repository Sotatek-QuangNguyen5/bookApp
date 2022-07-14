package services

import (
	"bookApp/dto"
	"bookApp/errs"
	"bookApp/repository"
	"bookApp/config"
)

type CategoryServices interface {

	ListCategory() ([]*dto.Category, *errs.AppError)
	CreateCategory(*dto.Category) (*errs.AppError)
	UpdateCategory(*dto.Category) (*errs.AppError)
	DeleteCategory(*dto.Category) (*errs.AppError)
	GetByIdCategory(*dto.Category) (*dto.Category, *errs.AppError)
	GetByBookId(*dto.Book) ([]*dto.Category, *errs.AppError)
}

type DefaultCategoryServices struct {

	repo repository.CategoryRepository
}

func NewCategoryServices(repo repository.CategoryRepository) CategoryServices {

	return DefaultCategoryServices{

		repo: repo,
	}
}

func (c DefaultCategoryServices) ListCategory() ([]*dto.Category, *errs.AppError) {

	res, err := c.repo.List()
	if err != nil {

		return nil, err
	}
	categories := dto.CategoriesModelToCategoriesDto(res)
	return categories, nil
}

func (c DefaultCategoryServices) DeleteCategory(category *dto.Category) (*errs.AppError) {

	err := dto.CheckID(category.Category_id)
	if err != nil {

		return err
	}
	return c.repo.Delete(category.Category_id)
}

func (c DefaultCategoryServices) UpdateCategory(category *dto.Category) (*errs.AppError) {

	err := dto.CheckID(category.Category_id)
	if err != nil {

		return err
	}
	return c.repo.Update(dto.CategoryDtoToCategoryModes(category))
}

func (c DefaultCategoryServices) CreateCategory(category *dto.Category) (*errs.AppError) {

	return c.repo.Create(dto.CategoryDtoToCategoryModes(category))
}

func (c DefaultCategoryServices) GetByIdCategory(category *dto.Category) (*dto.Category, *errs.AppError) {

	e := dto.CheckID(category.Category_id)
	if e != nil {

		return nil, e
	}
	res, err := c.repo.GetById(category.Category_id)
	if err != nil {

		return nil, err
	}
	categoryDto := dto.CategoryModelToCategoryDto(res)
	return categoryDto, nil
}

func (c DefaultCategoryServices) GetByBookId(book *dto.Book) ([]*dto.Category, *errs.AppError) {

	e := dto.CheckID(book.Book_id)
	if e != nil {

		return nil, e
	}
	b := NewBookServices(repository.NewBookRepository(config.DB))
	res, err := b.GetByIdBook(book)
	if err != nil {

		return nil, err
	}
	// if res.Categories == nil || len(res.Categories) == 0 {

	// 	return nil, errs.ErrorDataNotSurvive()
	// }
	return res.Categories, nil
}