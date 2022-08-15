package handlers

import (
	
	"bookApp/dto"
	"bookApp/errs"
	"bookApp/services"
	"net/http"
	"github.com/gin-gonic/gin"
)

type BookHandler struct {

	bookServices services.BookServices
}

func NewBookHandler(bookServices services.BookServices) BookHandler {

	return BookHandler{

		bookServices: bookServices,
	}
}

func (b BookHandler) GetListBook() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		author_id := ctx.MustGet("author_id")
		res, err := b.bookServices.ListBook(author_id.(int))

		if err != nil {

			WriteError(ctx, err)
			return
		}

		WriteRespon(ctx, http.StatusOK, res)
	}
}

func (b BookHandler) DeleteBook() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var book = new(dto.Book)
		err := ctx.ShouldBindJSON(book)
		if err != nil {

			WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}

		author_id := ctx.MustGet("author_id")
		e := b.bookServices.DeleteBook(book, author_id.(int))
		if e != nil {

			WriteError(ctx, e)
			return
		}

		WriteRespon(ctx, http.StatusOK, dto.MessageDeleteSuccess("Book"))
	}
}

func (b BookHandler) CreateBook() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var book = new(dto.Book)
		err := ctx.ShouldBindJSON(book)

		if err != nil {

			WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}

		e := b.bookServices.CreateBook(book)

		if e != nil {

			WriteError(ctx, e)
			return
		}

		WriteRespon(ctx, http.StatusOK, dto.MessageCreateSuccess("Book"))
	}
}

func (b BookHandler) UpdateBook() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var book = new(dto.Book)
		e := ctx.ShouldBindJSON(book)
		if e != nil {

			WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}

		author_id := ctx.MustGet("author_id")
		err := b.bookServices.UpdateBook(book, author_id.(int))

		if err != nil {

			WriteError(ctx, err)
			return
		}

		WriteRespon(ctx, http.StatusOK, dto.MessageUpdateSuccess("Book"))
	}
}

func (b BookHandler) GetByIdBook() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var book = new(dto.Book)
		e := ctx.ShouldBindJSON(book)
		if e != nil {

			WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}

		author_id := ctx.MustGet("author_id")
		res, err := b.bookServices.GetByIdBook(book, author_id.(int))
		if err != nil {

			WriteError(ctx, err)
			return
		}

		WriteRespon(ctx, http.StatusOK, res)
	}
}

func (b BookHandler) FilterBook() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var filter = new(dto.FilterBook)
		e := ctx.ShouldBindJSON(filter)
		if e != nil {

			WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}
		res, err := b.bookServices.FilterBook(filter)
		if err != nil {

			WriteError(ctx, err)
			return
		}
		WriteRespon(ctx, http.StatusOK, res)
	}
}

func (b BookHandler) AddCategory() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var bc = new(dto.BookCategory)
		e := ctx.ShouldBindJSON(bc)
		if e != nil {

			WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}
		err := b.bookServices.AddCategory(bc)
		if err != nil {

			WriteError(ctx, err)
			return
		}
		WriteRespon(ctx, http.StatusOK, dto.MessageAddSuccess("Category To Book"))
	}
}

func (b BookHandler) AddAuthor() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var ba = new(dto.BookAuthor)
		e := ctx.ShouldBindJSON(ba)
		if e != nil {

			WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}
		err := b.bookServices.AddAuthor(ba)
		if err != nil {

			WriteError(ctx, err)
			return
		}
		WriteRespon(ctx, http.StatusOK, dto.MessageAddSuccess("Author To Book"))
	}
}

func (b BookHandler) DeleteCategory() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var bc = new(dto.BookCategory)
		e := ctx.ShouldBindJSON(bc)
		if e != nil {

			WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}
		err := b.bookServices.DeleteCategory(bc)
		if err != nil {

			WriteError(ctx, err)
			return
		}
		WriteRespon(ctx, http.StatusOK, dto.MessageDeleteSuccess("Category From Book"))
	}
}

func (b BookHandler) DeleteAuthor() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var ba = new(dto.BookAuthor)
		e := ctx.ShouldBindJSON(ba)
		if e != nil {

			WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}
		err := b.bookServices.DeleteAuthor(ba)
		if err != nil {

			WriteError(ctx, err)
			return
		}
		WriteRespon(ctx, http.StatusOK, dto.MessageDeleteSuccess("Author From Book"))
	}
}