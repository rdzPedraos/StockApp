package resources

import (
	"app/integrations/financial"
	"app/models"
	"time"
)

type stockListItem struct {
	Ticker     string                        `json:"ticker"`
	Company    financial.CompanyDataResponse `json:"company"`
	LastReview time.Time                     `json:"last_review"`
	Sentiment  string                        `json:"sentiment"`
}

// StockListResource implementa la interfaz TypedResource para formatear listas de stocks
type StockListResource struct{}

// Format formatea un ticker específico
func (s StockListResource) Format(ticker models.Ticker) interface{} {
	var (
		lastReview time.Time
		sentiment  string
	)

	if len(ticker.Recommendations) > 0 {
		lastReview = ticker.Recommendations[0].Time
		sentiment = ticker.Recommendations[0].RatingTo.Label()
	}

	return stockListItem{
		Ticker:     ticker.ID,
		Company:    s.GetCompanyData(ticker),
		LastReview: lastReview,
		Sentiment:  sentiment,
	}
}

func (s StockListResource) GetCompanyData(ticker models.Ticker) financial.CompanyDataResponse {
	companyData, _ := financial.Service().GetCompanyData(ticker.ID)
	return companyData
}
