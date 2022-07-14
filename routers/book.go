package routers

import (
	"bookApp/config"
	"bookApp/handlers"
	"bookApp/repository"
	"bookApp/services"

	"github.com/gin-gonic/gin"
)

func BookRoute(router *gin.Engine) {


	h := handlers.NewBookHandler(services.NewBookServices(repository.NewBookRepository(config.DB)))
	r := router.Group("/book")
	{
		r.GET("/", h.GetListBook())
		r.PUT("/", h.UpdateBook())
		r.DELETE("/", h.DeleteBook())
		r.POST("/", h.CreateBook())
		r.GET("/detail", h.GetByIdBook())
		r.GET("/filter", h.FilterBook())
		r.DELETE("/delcate", h.DeleteCategory())
		r.POST("/addcate", h.AddCategory())
		r.DELETE("/delauthor", h.DeleteAuthor())
		r.POST("/addauthor", h.AddAuthor())
	}
}
