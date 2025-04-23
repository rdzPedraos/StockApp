package financial

import (
	"app/lib/cache"
	"log"
	"time"
)

type GetPriceResponse struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Volume float64 `json:"volume"`
}

func (s *service) GetPrice(ticker string) (float64, error) {
	var price float64

	err := cache.Once(
		"financial.GetPrice."+ticker,
		1*time.Hour,
		&price,
		func() (*float64, error) {
			log.Default().Println("Executing Financial.GetPrice method")
			endpoint := "/api/v3/quote-short/" + ticker + "?apikey=" + s.apiKey

			var apiResult []GetPriceResponse
			err := s.client.Get(endpoint, &apiResult)

			if err != nil {
				return nil, err
			}

			log.Default().Println("Financial.GetPrice method executed successfully")
			value := apiResult[0].Price
			return &value, nil
		})

	return price, err
}
