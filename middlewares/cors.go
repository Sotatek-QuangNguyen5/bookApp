package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {

	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "POST", "DELETE", "GET"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	})
}
