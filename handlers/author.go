package handlers

import (
	"bookApp/dto"
	"bookApp/errs"
	"bookApp/services"
	"net/http"

	"github.com/gin-gonic/gin"
)


type AuthorHandler struct {

	services services.AuthorServices
}


func NewAuthorHandler(services services.AuthorServices) AuthorHandler {

	return AuthorHandler{

		services: services,
	}
}

func (a AuthorHandler) Login() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var author = new(dto.Author)
		e := ctx.ShouldBindJSON(author)
		if e != nil {

			WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}
		res, err := a.services.Login(author)
		if err != nil {

			WriteError(ctx, err)
			return
		}
		WriteRespon(ctx, http.StatusOK, res)
	}
}



func (a AuthorHandler) GetListAuthor() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		res, err := a.services.ListAuthor()
		if err != nil {

			WriteError(ctx, err)
			return
		}
		WriteRespon(ctx, http.StatusOK, res)
	}
}

func (a AuthorHandler) DeleteAuthor() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var author = new(dto.Author)
		err := ctx.ShouldBindJSON(author)
		if err != nil {

			WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}
		e := a.services.DeleteAuthor(author)
		if e != nil {

			WriteError(ctx, e)
			return
		}
		WriteRespon(ctx, http.StatusOK, dto.MessageDeleteSuccess("Author"))
	}
}

func (a AuthorHandler) CreateAuthor() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var author = new(dto.Author)
		err := ctx.ShouldBindJSON(author)
		if err != nil {

			WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}
		e := a.services.CreateAuthor(author)
		if e != nil {

			WriteError(ctx, e)
			return
		}
		WriteRespon(ctx, http.StatusOK, dto.MessageCreateSuccess("Author"))
	}
}

func (a AuthorHandler) UpdateAuthor() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var author = new(dto.Author)
		err := ctx.ShouldBindJSON(author)
		if err != nil {

			WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}
		e := a.services.UpdateAuthor(author)
		if e != nil {

			WriteError(ctx, e)
			return
		}
		WriteRespon(ctx, http.StatusOK, dto.MessageUpdateSuccess("Author"))
	}
}

func (a AuthorHandler) GetByIdAuthor() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var author = new(dto.Author)
		err := ctx.ShouldBindJSON(author)
		if err != nil {

			WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}
		res, e := a.services.GetByIdAuthor(author)
		if e != nil {

			WriteError(ctx, e)
			return
		}

		WriteRespon(ctx, http.StatusOK, res)
	}
}

func (a AuthorHandler) GetByBookId() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var book = new(dto.Book)
		e := ctx.ShouldBindJSON(book)
		if e != nil {

			WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}
		author_id := ctx.MustGet("author_id")
		res, err := a.services.GetByBookId(book, author_id.(int))
		if err != nil {

			WriteError(ctx, err)
			return
		}
		WriteRespon(ctx, http.StatusOK, res)
	}
}