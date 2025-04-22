package resources

import (
	"app/models"
	"time"
)

type stockRecommendation struct {
	Time           time.Time `json:"time"`
	Recommendation string    `json:"recommendation"`
}

type stockListItem struct {
	Ticker             string              `json:"ticker"`
	Company            string              `json:"company"`
	Price              float64             `json:"price"`
	Logo               string              `json:"logo"`
	MarketCap          float64             `json:"market_cap"`
	LastRecommendation stockRecommendation `json:"last_recommendation"`
}

// StockListResource implementa la interfaz TypedResource para formatear listas de stocks
type StockListResource struct{}

// FormatTyped formatea un ticker específico
func (s StockListResource) Format(ticker models.Ticker) interface{} {
	lastRecommendation := stockRecommendation{}

	// Verificar si hay recomendaciones disponibles
	if len(ticker.Recommendations) > 0 {
		lastRecommendation = stockRecommendation{
			Time:           ticker.Recommendations[0].Time,
			Recommendation: ticker.Recommendations[0].RatingTo.Label(),
		}
	}

	return stockListItem{
		Ticker:             ticker.ID,
		Company:            ticker.Company,
		Price:              0, // Actualizar con datos reales si están disponibles
		Logo:               "https://s2.coinmarketcap.com/static/img/coins/64x64/1027.png",
		MarketCap:          0, // Actualizar con datos reales si están disponibles
		LastRecommendation: lastRecommendation,
	}
}
