package resources

import (
	"app/config"
	"app/integrations/financial"
	"app/models"
	"log"
	"time"
)

type stockListItem struct {
	Ticker     string    `json:"ticker"`
	Company    string    `json:"company"`
	Price      float64   `json:"price"`
	MarketCap  float64   `json:"market_cap"`
	Logo       string    `json:"logo"`
	LastReview time.Time `json:"last_review"`
	Sentiment  string    `json:"sentiment"`
}

// StockListResource implementa la interfaz TypedResource para formatear listas de stocks
type StockListResource struct{}

// Format formatea un ticker específico
func (s StockListResource) Format(ticker models.Ticker) interface{} {
	var (
		lastReview time.Time
		sentiment  string
	)

	if len(ticker.Recommendations) > 0 {
		lastReview = ticker.Recommendations[0].Time
		sentiment = ticker.Recommendations[0].RatingTo.Label()
	}

	price := make(chan float64, 1)
	marketCap := make(chan float64, 1)

	go s.GetPrice(ticker, price)
	go s.GetMarketCap(ticker, marketCap)

	return stockListItem{
		Ticker:     ticker.ID,
		Company:    ticker.Company,
		Price:      <-price,
		MarketCap:  <-marketCap,
		LastReview: lastReview,
		Sentiment:  sentiment,
		Logo:       config.Server.Url + "/api/tickers/" + ticker.ID + "/logo",
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
