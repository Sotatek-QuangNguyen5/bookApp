package dto

import "bookApp/models"

type Category struct {
	Category_id int    `json:"category_id"`
	Name        string `json:"name"`
	Decription  string `json:"decription"`
}

func CategoryModelToCategoryDto(category *models.Category) *Category {

	if category == nil {

		return nil
	}
	return &Category{

		Category_id: category.Category_id,
		Name: category.Name,
		Decription: category.Decription,
	}
}

func CategoriesModelToCategoriesDto(categoriesModel []*models.Category) []*Category {

	var categories []*Category
	for _, category := range categoriesModel {

		categories = append(categories, CategoryModelToCategoryDto(category))
	}
	return categories
}

func CategoryDtoToCategoryModes(category *Category) *models.Category {

	if category == nil {

		return nil
	}
	return &models.Category{

		Category_id: category.Category_id,
		Name: category.Name,
		Decription: category.Decription,
	}
}

func CategoriesDtoToCategoriesModel(categoriesDto []*Category) []*models.Category {

	var categories []*models.Category
	for _, category := range categoriesDto {

		categories = append(categories, CategoryDtoToCategoryModes(category))
	}
	return categories
}