package stockRating

import (
	"log"
	"time"
)

type getAllItem struct {
	Ticker     string    `json:"ticker"`
	TargetFrom string    `json:"target_from"`
	TargetTo   string    `json:"target_to"`
	Company    string    `json:"company"`
	Action     string    `json:"action"`
	Brokerage  string    `json:"brokerage"`
	RatingFrom string    `json:"rating_from"`
	RatingTo   string    `json:"rating_to"`
	Time       time.Time `json:"time"`
}

type getAllResponse struct {
	Items    []getAllItem `json:"items"`
	NextPage string       `json:"next_page"`
}

func (s *service) GetAll() ([]getAllItem, error) {
	log.Default().Println("Executing StockRating.GetAll method")

	allItems := []getAllItem{}
	endpoint := "/swechallenge/list"
	queryParams := map[string]string{
		"next_page": "",
	}

	for {
		var response getAllResponse
		err := s.client.GetWithQueryParams(endpoint, &response, queryParams)

		if err != nil {
			return nil, err
		}

		// Agregar los ratings a la colección total
		allItems = append(allItems, response.Items...)

		if response.NextPage == "" {
			break
		}

		queryParams["next_page"] = response.NextPage
	}

	log.Default().Println("StockRating.GetAll method executed successfully")
	return allItems, nil
}
