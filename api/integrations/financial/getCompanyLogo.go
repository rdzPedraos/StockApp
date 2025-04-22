package financial

import (
	"app/lib/apiClient"
	"log"
)

func (s *service) GetCompanyLogo(ticker string) (*[]byte, string, error) {
	log.Default().Println("Executing Financial.GetCompanyLogo method")
	endpoint := "/image-stock/" + ticker + ".png?apikey=" + s.apiKey

	var result apiClient.BinaryResult
	err := s.client.GetBinary(endpoint, &result)

	if err != nil {
		return nil, "", err
	}

	log.Default().Println("Financial.GetCompanyLogo method executed successfully")
	return &result.Data, result.ContentType, nil
}
