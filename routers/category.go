package routers

import (
	"bookApp/config"
	"bookApp/handlers"
	"bookApp/repository"
	"bookApp/services"

	"github.com/gin-gonic/gin"
)

func CategoryRoute(router *gin.Engine) {

	h := handlers.NewCategoryHandler(services.NewCategoryServices(repository.NewCategoryReponsitory(config.DB)))
	r := router.Group("/category")
	{

		r.GET("/", h.GetListCategory())
		r.POST("/", h.CreateCategory())
		r.PUT("/", h.UpdateCategory())
		r.DELETE("/", h.DeleteCategory())
		r.GET("/detail", h.GetByIdCategory())
		r.GET("/getbybookid", h.GetByBookId())
	}
}