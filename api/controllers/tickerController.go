package controllers

import (
	"app/database/connection"
	"app/database/scopes"
	"app/lib/helper"
	"app/models"
	"app/resources"
	"math"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TickerController struct {
}

func (c *TickerController) List(ctx *gin.Context) {
	page := helper.ClampString(ctx.Query("page"), 1, 100)
	perPage := helper.ClampString(ctx.Query("per_page"), 10, 100)

	db := connection.DB

	var results []models.Ticker
	var total int64

	db.Model(&models.Ticker{}).Count(&total)

	db.Scopes(scopes.Pagination(page, perPage)).
		Preload("Recommendations", func(db *gorm.DB) *gorm.DB {
			return db.Order("recommendations.time DESC").Limit(1)
		}).
		Find(&results)

	formattedResults := resources.FormatArrayWith(resources.StockListResource{}, results)

	lastPage := math.Ceil(float64(total) / float64(perPage))

	ctx.JSON(200, gin.H{
		"data": formattedResults,
		"meta": gin.H{
			"total":     total,
			"page":      page,
			"per_page":  perPage,
			"last_page": lastPage,
		},
	})
}

func (c *TickerController) Show(ctx *gin.Context) {

}
