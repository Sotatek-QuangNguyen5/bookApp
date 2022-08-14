package routers

import (
	
	"bookApp/config"
	"bookApp/middlewares"
	"log"
	"github.com/gin-gonic/gin"
)

func Start() {

	config.InitDatabase()
	router := gin.Default()
	router.Use(middlewares.Cors())
	router.POST("/token", middlewares.GenerateToken())

	Token(router)
	router.Use(middlewares.VerifyToken())
	{
		BookRoute(router)
		AuthorRoute(router)
		CategoryRoute(router)
		ConversationRoute(router)
	}
	log.Println("Server is running on PORT ", config.GetPort())
	router.Run(config.GetPort())
}
