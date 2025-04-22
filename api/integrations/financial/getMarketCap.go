package financial

import (
	"log"
)

type MarketCapResponse struct {
	Symbol    string  `json:"symbol"`
	Date      string  `json:"date"`
	MarketCap float64 `json:"marketCap"`
}

func (s *service) GetMarketCap(ticker string) (float64, error) {
	log.Default().Println("Executing Financial.GetMarketCap method")
	endpoint := "/api/v3/market-capitalization/" + ticker + "?apikey=" + s.apiKey

	var result []MarketCapResponse
	err := s.client.Get(endpoint, &result)

	if err != nil {
		return 0, err
	}

	log.Default().Println("Financial.GetMarketCap method executed successfully")
	return result[0].MarketCap, nil
}
