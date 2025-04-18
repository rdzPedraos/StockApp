package main

import (
	"app/cmd"
	"app/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// only if any command not was executed, run the server
	if cmd.Execute() {
		return
	}

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port))
}
