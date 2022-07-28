package repository

import (
	"bookApp/errs"
	"bookApp/models"
	"database/sql"
	"fmt"
)

type CategoryRepository interface {

	List() ([]*models.Category, *errs.AppError)
	Delete(int) (*errs.AppError)
	Update(*models.Category) (*errs.AppError)
	Create(*models.Category) (*errs.AppError)
	GetById(int) (*models.Category, *errs.AppError)
	GetByBookId(int) ([]*models.Category, *errs.AppError)
}

type DefaultCategoryRepository struct {

	db *sql.DB
}

func NewCategoryReponsitory(db *sql.DB) CategoryRepository {

	return DefaultCategoryRepository{

		db: db,
	}
}

func (c DefaultCategoryRepository) List() ([]*models.Category, *errs.AppError) {

	res, err := c.db.Query("SELECT * FROM category")
	if err != nil {

		return nil, errs.ErrorGetData()
	}

	var categories []*models.Category
	for res.Next() {

		var category = new(models.Category)
		e := res.Scan(&category.Category_id, &category.Name, &category.Decription)
		if e != nil {

			return nil, errs.ErrorReadData()
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (c DefaultCategoryRepository) Delete(Category_id int) (*errs.AppError) {

	_, e := c.GetById(Category_id)
	if e != nil {

		return e
	}
	query := fmt.Sprintf("DELETE FROM book_category WHERE category_id = %d", Category_id)
	_, err := c.db.Query(query)
	if err != nil {

		return errs.ErrorDeleteData()
	}

	query = fmt.Sprintf("DELETE FROM category WHERE category_id = %d", Category_id)
	_, err = c.db.Query(query)
	if err != nil {

		return errs.ErrorDeleteData()
	}

	return nil
}

func (c DefaultCategoryRepository) Create(category *models.Category) (*errs.AppError) {

	query := fmt.Sprintf("INSERT INTO category(name, description) VALUES('%s','%s')", category.Name, category.Decription)
	_, err := c.db.Query(query)
	if err != nil {

		return errs.ErrorInsertData()
	}
	return nil
}

func (c DefaultCategoryRepository) Update(category *models.Category) (*errs.AppError) {

	_, e := c.GetById(category.Category_id)
	if e != nil {

		return e
	}
	query := fmt.Sprintf("UPDATE category SET name = '%s', description = '%s' WHERE category_id = %d", category.Name, category.Decription, category.Category_id)
	_, err := c.db.Query(query)
	if err != nil {

		return errs.ErrorUpdateData()
	}
	return nil
}

func (c DefaultCategoryRepository) GetById(category_id int) (*models.Category, *errs.AppError) {

	query := fmt.Sprintf("SELECT * FROM category WHERE category_id = %d", category_id)
	res, e := c.db.Query(query)
	if e != nil {

		return nil, errs.ErrorGetData()
	}

	var category *models.Category
	for res.Next() {

		err := res.Scan(&category.Category_id, &category.Name, &category.Decription)
		if err != nil {

			return nil, errs.ErrorReadData()
		}
	}
	if category.Category_id == 0 {

		return nil, errs.ErrorDataNotSurvive()
	}

	return category, nil
}

func (c DefaultCategoryRepository) GetByBookId(book_id int) ([]*models.Category, *errs.AppError) {

	query := fmt.Sprintf("SELECT c.* FROM category c JOIN book_category bc ON c.category_id = bc.category_id AND bc.book_id = %d", book_id)
	res, e := c.db.Query(query)
	if e != nil {

		return nil, errs.ErrorGetData()
	}

	var categories []*models.Category
	for res.Next() {

		var category = new(models.Category)
		err := res.Scan(&category.Category_id, &category.Name, &category.Decription)
		if err != nil {

			return nil, errs.ErrorReadData()
		}
		categories = append(categories, category)
	}

	return categories, nil
}