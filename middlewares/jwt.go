package middlewares

import (
	
	"bookApp/config"
	"bookApp/dto"
	"bookApp/errs"
	"bookApp/handlers"
	"bookApp/repository"
	"bookApp/services"
	"bookApp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyToken() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("Authorization")
		if token == "" {

			handlers.WriteError(ctx, errs.NewUnauthenticatedError("Not Authorization"))
			return
		}

		res, err := config.VerifyToken(token)
		if err != nil {

			handlers.WriteError(ctx, err)
			return
		}
		var author = new(dto.Author)
		err = utils.MapToStruct(res["data"], author)
		if err != nil {

			handlers.WriteError(ctx, err)
			return
		}
		serviceAuthor := services.NewAuthorServices(repository.NewAuthorRepository(config.DB))
		_, err = serviceAuthor.GetByIdAuthor(author)
		if err != nil {

			handlers.WriteError(ctx, err)
			return
		}
		ctx.Set("author_id", author.Author_id)
		ctx.Next()
	}
}

func GenerateToken() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var req map[string]interface{}
		err := ctx.ShouldBindJSON(&req)
		if err != nil {

			handlers.WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}
		res, e := config.NewAccessJsonWebToken(req)
		if e != nil {

			handlers.WriteError(ctx, e)
			return
		}
		handlers.WriteRespon(ctx, http.StatusOK, res)
	}
}

func RefreshToken() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		refreshToken := ctx.Request.Header.Get("RefreshAuthorization")
		if refreshToken == "" {

			handlers.WriteError(ctx, errs.NewUnauthenticatedError("Not authorization!!!"))
			return
		}
		res, err := config.VerifyToken(refreshToken)
		if err != nil {

			handlers.WriteError(ctx, err)
			return
		}
		var author = new(dto.Author)
		err = utils.MapToStruct(res["data"], author)
		if err != nil {

			handlers.WriteError(ctx, err)
			return
		}
		serviceAuthor := services.NewAuthorServices(repository.NewAuthorRepository(config.DB))
		respon, err := serviceAuthor.GetByIdAuthor(author)
		if err != nil {

			handlers.WriteError(ctx, err)
			return
		}
		accessToken, err := config.NewAccessJsonWebToken(utils.StructToMap(respon))
		if err != nil {

			handlers.WriteError(ctx, err)
			return
		}
		authorization := &dto.Authentication{

			RefreshToken: refreshToken,
			AccessToken: *accessToken,
			Message: "Refresh Token OK!!!",
		}
		handlers.WriteRespon(ctx, http.StatusOK, authorization)
	}
}