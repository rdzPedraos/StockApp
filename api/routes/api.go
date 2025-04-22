package routes

import (
	"app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterApiRoutes(router *gin.RouterGroup) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	tickerRoutes(router.Group("/tickers"))
}

func tickerRoutes(group *gin.RouterGroup) {
	tickerController := controllers.TickerController{}

	group.GET("", tickerController.List)
	group.GET("/:id", tickerController.Show)
	group.GET("/:id/logo", tickerController.GetCompanyLogo)
}
