package main

import (
	"app/config"
	"app/db"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Get()
	cfg := config.Get()

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port))
}
