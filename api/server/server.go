package server

import (
	"app/config"
	"app/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	address := fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port)

	routes.RegisterApiRoutes(router.Group("/api"))

	router.Run(address)
	log.Printf("Server running on %s", address)
}
