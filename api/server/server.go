package server

import (
	"app/config"
	"app/routes"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	address := fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port)

	router.Use(corsMiddleware())
	registerRoutes(router)

	router.Run(address)
	log.Printf("Server running on %s", address)
}

func corsMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})
}

func registerRoutes(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routes.RegisterV1Routes(router.Group("/"))
}
