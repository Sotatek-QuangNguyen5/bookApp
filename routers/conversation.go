package routers

import (
	
	"bookApp/config"
	"bookApp/handlers"
	"bookApp/repository"
	"bookApp/services"

	"github.com/gin-gonic/gin"
)

func ConversationRoute(router *gin.Engine) {

	h := handlers.NewAuthorHandler(services.NewAuthorServices(repository.NewAuthorRepository(config.DB)))
	c := handlers.NewConversationHandler(services.NewConversationServices(repository.NewConversationRepository(config.DB)))

	router.GET("/message", h.GetListAuthor())
	r := router.Group("/message")
	{

		r.GET("/:conversation_id", c.GetListMess())
	}
}