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

	BookRoute(router)
	AuthorRoute(router)
	CategoryRoute(router)

	log.Println("Server is running on PORT ", config.PORT)
	router.Run(config.PORT)
}
