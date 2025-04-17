package stockRating

import (
	"fmt"
	"net/url"
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
	items    []getAllItem
	nextPage string
}

func (s *service) GetAll() ([]getAllItem, error) {
	allItems := []getAllItem{}
	endpoint := "/swechallenge/list"
	nextPage := ""

	for {
		requestURL := endpoint

		if nextPage != "" {
			encodedNextPage := url.QueryEscape(nextPage)
			requestURL = fmt.Sprintf("%s?next_page=%s", endpoint, encodedNextPage)
		}

		var response getAllResponse
		err := s.client.Get(requestURL, &response)
		if err != nil {
			return nil, fmt.Errorf("error al obtener datos de ratings de acciones: %w", err)
		}

		// Agregar los ratings a la colección total
		allItems = append(allItems, response.items...)

		if response.nextPage == "" {
			break
		}

		nextPage = response.nextPage
	}

	return allItems, nil
}
