package financial

import (
	"app/lib/apiClient"
	"app/lib/cache"
	"log"
	"time"
)

// LogoResult estructura para almacenar los datos de logo y su tipo de contenido
type LogoResult struct {
	Data        []byte `json:"data"`
	ContentType string `json:"contentType"`
}

func (s *service) GetCompanyLogo(ticker string) (*[]byte, string, error) {
	var logoResult LogoResult

	err := cache.Once(
		"financial.GetCompanyLogo."+ticker,
		24*31*time.Hour,
		&logoResult,
		func() (*LogoResult, error) {
			log.Default().Println("Executing Financial.GetCompanyLogo method")
			endpoint := "/image-stock/" + ticker + ".png?apikey=" + s.apiKey

			var apiResult apiClient.BinaryResult
			err := s.client.GetBinary(endpoint, &apiResult)

			if err != nil {
				return nil, err
			}

			log.Default().Println("Financial.GetCompanyLogo method executed successfully")
			return &LogoResult{
				Data:        apiResult.Data,
				ContentType: apiResult.ContentType,
			}, nil
		})

	if err != nil {
		return nil, "", err
	}

	return &logoResult.Data, logoResult.ContentType, nil
}
