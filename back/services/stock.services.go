package services

import (
	"app/config"
	"app/utils"
	"fmt"
	"log"
	"net/url"
	"time"
)

type StockItem struct {
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

// StockRatingResponse representa la respuesta de la API con paginación
type StockItemResponse struct {
	Items    []StockItem `json:"items"`
	NextPage string      `json:"next_page"`
}

// StockRatingService proporciona métodos para interactuar con la API de ratings de acciones
type StockRatingService struct {
	apiClient *utils.APIClient
}

func Init() *StockRatingService {
	baseURL := config.Get().Services.Stock.API_URL
	apiKey := config.Get().Services.Stock.API_KEY

	if baseURL == "" || apiKey == "" {
		log.Fatal("API URL or API Key is not set")
	}

	apiClient := utils.NewAPIClient(baseURL)
	apiClient.SetAuthToken(apiKey)

	return &StockRatingService{
		apiClient: apiClient,
	}
}

func (s *StockRatingService) GetAllItems() ([]StockItem, error) {
	allItems := []StockItem{}
	endpoint := "/swechallenge/list"
	nextPage := ""

	for {
		// Construir la URL con el parámetro de next_page si es necesario
		requestURL := endpoint
		if nextPage != "" {
			// Escapar el valor de next_page para asegurar que se codifique correctamente
			encodedNextPage := url.QueryEscape(nextPage)
			requestURL = fmt.Sprintf("%s?next_page=%s", endpoint, encodedNextPage)
		}

		// Hacer la petición a la API
		var response StockItemResponse
		err := s.apiClient.Get(requestURL, &response)
		if err != nil {
			return nil, fmt.Errorf("error al obtener datos de ratings de acciones: %w", err)
		}

		// Agregar los ratings a la colección total
		allItems = append(allItems, response.Items...)

		// Verificar si hay más páginas
		if response.NextPage == "" {
			// No hay más páginas, terminamos
			break
		}

		// Actualizar el token para la siguiente página
		nextPage = response.NextPage
	}

	return allItems, nil
}
