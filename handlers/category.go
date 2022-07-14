package handlers

import (

	"bookApp/dto"
	"bookApp/errs"
	"bookApp/services"
	"net/http"
	"github.com/gin-gonic/gin"
)

type CategoryHander struct {

	categoryServices services.CategoryServices
}

func NewCategoryHandler(categoryServices services.CategoryServices) CategoryHander {

	return CategoryHander{

		categoryServices: categoryServices,
	}
}

func (c CategoryHander) GetListCategory() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		res, err := c.categoryServices.ListCategory()
		if err != nil {

			WriteError(ctx, err)
			return
		}
		WriteRespon(ctx, http.StatusOK, res)
	}
}

func (c CategoryHander) DeleteCategory() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var category = new(dto.Category)
		err := ctx.ShouldBindJSON(category)
		if err != nil {

			WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}
		WriteRespon(ctx, http.StatusOK, dto.MessageDeleteSuccess("Category"))
	}
}

func (c CategoryHander) CreateCategory() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var category = new(dto.Category)
		err := ctx.ShouldBindJSON(category)
		if err != nil {

			WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}
		e := c.categoryServices.CreateCategory(category)
		if e != nil {

			WriteError(ctx, e)
			return
		}
		WriteRespon(ctx, http.StatusOK, dto.MessageCreateSuccess("Category"))
	}
}

func (c CategoryHander) UpdateCategory() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var category = new(dto.Category)
		err := ctx.ShouldBindJSON(category)
		if err != nil {

			WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}
		WriteRespon(ctx, http.StatusOK, dto.MessageUpdateSuccess("Category"))
	}
}

func (c CategoryHander) GetByIdCategory() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var category = new(dto.Category)
		e := ctx.ShouldBindJSON(category)
		if e != nil {

			WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}
		res, err := c.categoryServices.GetByIdCategory(category)
		if err != nil {

			WriteError(ctx, err)
			return
		}
		WriteRespon(ctx, http.StatusOK, res)
	}
}

func (c CategoryHander) GetByBookId() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var book = new(dto.Book)
		e := ctx.ShouldBindJSON(book)
		if e != nil {

			WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}
		res, err := c.categoryServices.GetByBookId(book)
		if err != nil {

			WriteError(ctx, err)
			return
		}
		WriteRespon(ctx, http.StatusOK, res)
	}
}