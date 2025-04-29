package routes

import (
	"app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterV1Routes(router *gin.RouterGroup) {
	tickerRoutes(router.Group("/tickers"))
}

func tickerRoutes(group *gin.RouterGroup) {
	tickerController := controllers.TickerController{}

	group.GET("", tickerController.List)
	group.GET("/:id", tickerController.Show)
}
