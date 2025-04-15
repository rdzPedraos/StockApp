package utils

import (
	"app/models"
)

var Schemas = []interface{}{
	new(models.Broker),
	new(models.Ticker),
	new(models.Recommendation),
}
