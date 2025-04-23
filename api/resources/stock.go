package resources

import (
	"app/config"
	"app/integrations/financial"
	"app/models"
)

type stockItem struct {
	Ticker          string        `json:"ticker"`
	Company         string        `json:"company"`
	Price           float64       `json:"price"`
	MarketCap       float64       `json:"market_cap"`
	Logo            string        `json:"logo"`
	Recommendations []interface{} `json:"recommendations"`
}

// StockListResource implementa la interfaz TypedResource para formatear listas de stocks
type StockResource struct{}

// Format formatea un ticker específico
func (s StockResource) Format(ticker models.Ticker) interface{} {
	price := make(chan float64, 1)
	marketCap := make(chan float64, 1)

	go s.GetPrice(ticker, price)
	go s.GetMarketCap(ticker, marketCap)

	return stockItem{
		Ticker:          ticker.ID,
		Company:         ticker.Company,
		Price:           <-price,
		MarketCap:       <-marketCap,
		Logo:            config.Server.Url + "/api/tickers/" + ticker.ID + "/logo",
		Recommendations: FormatArrayWith(RecommendationResource{}, ticker.Recommendations),
	}
}

func (s StockResource) GetPrice(ticker models.Ticker, ch chan float64) {
	price, _ := financial.Service().GetPrice(ticker.ID)
	ch <- price
}

func (s StockResource) GetMarketCap(ticker models.Ticker, ch chan float64) {
	marketCap, _ := financial.Service().GetMarketCap(ticker.ID)
	ch <- marketCap
}
