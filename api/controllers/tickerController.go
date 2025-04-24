package controllers

import (
	"app/database/connection"
	"app/lib/helper"
	"app/lib/paginate"
	"app/models"
	"app/resources"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TickerController struct {
}

func (c *TickerController) List(ctx *gin.Context) {
	page := helper.ClampString(ctx.Query("page"), 1, 100)
	perPage := helper.ClampString(ctx.Query("per_page"), 10, 100)

	query := func(db *gorm.DB) *gorm.DB {
		return db.Preload("Recommendations", func(db *gorm.DB) *gorm.DB {
			return db.Order("recommendations.time DESC")
		})
	}

	response := new(paginate.PaginateResolver[models.Ticker]).
		Stmt(query).
		Resource(resources.StockListResource{}).
		Paginate(page, perPage).
		Resolve()

	ctx.JSON(200, response)
}

func (c *TickerController) Show(ctx *gin.Context) {
	tickerId := ctx.Param("id")

	if tickerId == "" {
		ctx.JSON(400, gin.H{"error": "Ticker ID is required"})
		return
	}

	ticker := models.Ticker{ID: tickerId}
	connection.DB.Preload("Recommendations", func(db *gorm.DB) *gorm.DB {
		return db.Order("recommendations.time DESC")
	}).Preload("Recommendations.Broker").Find(&ticker)

	if ticker.Company == "" {
		ctx.JSON(404, gin.H{"error": "Ticker not found"})
		return
	}

	resource := resources.FormatWith(resources.StockResource{}, ticker)
	ctx.JSON(200, resource)
}
