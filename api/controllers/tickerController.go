package controllers

import (
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
			return db.Order("recommendations.time DESC").Limit(1)
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

}
