package main

import (
	"app/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port))
}
