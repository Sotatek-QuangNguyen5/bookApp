package routers

import (
	"bookApp/config"
	"bookApp/handlers"
	"bookApp/repository"
	"bookApp/services"
	"github.com/gin-gonic/gin"
)

func AuthorRoute(router *gin.Engine) {

	h := handlers.NewAuthorHandler(services.NewAuthorServices(repository.NewAuthorRepository(config.DB)))
	r := router.Group("/author")
	{

		r.GET("/", h.GetListAuthor())
		r.DELETE("/", h.DeleteAuthor())
		r.PUT("/", h.UpdateAuthor())
		r.GET("/detail", h.GetByIdAuthor())
		r.GET("/getbybookid", h.GetByBookId())
	}
}