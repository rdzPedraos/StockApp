package financial

import (
	"log"
)

type CurrentPriceResponse struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Volume float64 `json:"volume"`
}

func (s *service) GetPrice(ticker string) (float64, error) {
	log.Default().Println("Executing Financial.GetPrice method")
	endpoint := "/api/v3/quote-short/" + ticker + "?apikey=" + s.apiKey

	var result []CurrentPriceResponse
	err := s.client.Get(endpoint, &result)

	if err != nil {
		return 0, err
	}

	log.Default().Println("Financial.GetPrice method executed successfully")
	return result[0].Price, nil
}
