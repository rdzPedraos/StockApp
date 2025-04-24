package financial

import (
	"app/lib/cache"
	"log"
	"time"
)

type CompanyDataResponse struct {
	Name          string  `json:"companyName"`
	Price         float64 `json:"price"`
	MarketCap     float64 `json:"marketCap"`
	Beta          float64 `json:"beta"`
	Change        float64 `json:"change"`
	ChangePercent float64 `json:"changePercentage"`
	Volume        float64 `json:"volume"`
	Exchange      string  `json:"exchange"`
	Sector        string  `json:"sector"`
	Industry      string  `json:"industry"`
	Website       string  `json:"website"`
	Image         string  `json:"image"`
}

func (s *service) GetCompanyData(ticker string) (CompanyDataResponse, error) {
	var companyData CompanyDataResponse

	err := cache.Once(
		"financial.GetCompanyData."+ticker,
		1*time.Hour,
		&companyData,
		func() (*CompanyDataResponse, error) {
			log.Default().Println("Executing Financial.GetCompanyData method")
			endpoint := "/stable/profile?symbol=" + ticker + "&apikey=" + s.apiKey

			var apiResult []CompanyDataResponse
			err := s.client.Get(endpoint, &apiResult)

			if err != nil {
				log.Default().Println("Financial.GetCompanyData method failed", err)
				return nil, err
			}

			if len(apiResult) == 0 {
				log.Default().Println("Financial.GetCompanyData method failed: no data returned")
				return nil, nil
			}

			log.Default().Println("Financial.GetCompanyData method executed successfully")
			return &apiResult[0], nil
		})

	return companyData, err
}
