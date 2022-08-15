package routers

import (
	"bookApp/config"
	"bookApp/handlers"
	"bookApp/middlewares"
	"bookApp/repository"
	"bookApp/services"

	"github.com/gin-gonic/gin"
)

func Token(router *gin.Engine) {

	handlerAuthor := handlers.NewAuthorHandler(services.NewAuthorServices(repository.NewAuthorRepository(config.DB)))
	router.POST("/login", handlerAuthor.Login())
	router.POST("/register", handlerAuthor.CreateAuthor())
	router.POST("/refreshtoken", middlewares.RefreshToken())
}