package resources

import (
	"app/config"
	"app/integrations/financial"
	"app/models"
	"log"
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
	MarketCap          float64             `json:"market_cap"`
	LastRecommendation stockRecommendation `json:"last_recommendation"`
	Logo               string              `json:"logo"`
}

// StockListResource implementa la interfaz TypedResource para formatear listas de stocks
type StockListResource struct{}

// Format formatea un ticker específico
func (s StockListResource) Format(ticker models.Ticker) interface{} {
	// Preparar la recomendación
	lastRecommendation := stockRecommendation{}
	if len(ticker.Recommendations) > 0 {
		lastRecommendation = stockRecommendation{
			Time:           ticker.Recommendations[0].Time,
			Recommendation: ticker.Recommendations[0].RatingTo.Label(),
		}
	}

	price := make(chan float64, 1)
	marketCap := make(chan float64, 1)

	go s.GetPrice(ticker, price)
	go s.GetMarketCap(ticker, marketCap)

	return stockListItem{
		Ticker:             ticker.ID,
		Company:            ticker.Company,
		Price:              <-price,
		MarketCap:          <-marketCap,
		LastRecommendation: lastRecommendation,
		Logo:               config.Server.Url + "/api/tickers/" + ticker.ID + "/logo",
	}
}

func (s StockListResource) GetPrice(ticker models.Ticker, ch chan float64) {
	price, err := financial.Service().GetPrice(ticker.ID)

	if err != nil {
		log.Println(err)
	}

	ch <- price
}

func (s StockListResource) GetMarketCap(ticker models.Ticker, ch chan float64) {
	marketCap, err := financial.Service().GetMarketCap(ticker.ID)

	if err != nil {
		log.Println(err)
	}

	ch <- marketCap
}
