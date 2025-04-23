package financial

import (
	"app/lib/cache"
	"log"
	"time"
)

type MarketCapResponse struct {
	Symbol    string  `json:"symbol"`
	Date      string  `json:"date"`
	MarketCap float64 `json:"marketCap"`
}

func (s *service) GetMarketCap(ticker string) (float64, error) {
	var marketCap float64

	err := cache.Once(
		"financial.GetMarketCap."+ticker,
		1*time.Hour,
		&marketCap,
		func() (*float64, error) {
			log.Default().Println("Executing Financial.GetMarketCap method")
			endpoint := "/api/v3/market-capitalization/" + ticker + "?apikey=" + s.apiKey

			var apiResult []MarketCapResponse
			err := s.client.Get(endpoint, &apiResult)

			if err != nil {
				return nil, err
			}

			log.Default().Println("Financial.GetMarketCap method executed successfully")
			value := apiResult[0].MarketCap
			return &value, nil
		})

	return marketCap, err
}
