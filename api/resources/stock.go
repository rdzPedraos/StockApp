package resources

import (
	"app/integrations/financial"
	"app/integrations/gemini"
	"app/models"
)

type stockItem struct {
	Ticker          string                        `json:"ticker"`
	Recommendations []interface{}                 `json:"recommendations"`
	Company         financial.CompanyDataResponse `json:"company"`
	Advice          string                        `json:"advice"`
}

// StockListResource implementa la interfaz TypedResource para formatear listas de stocks
type StockResource struct{}

// Format formatea un ticker específico
func (s StockResource) Format(ticker models.Ticker) interface{} {
	companyData := make(chan financial.CompanyDataResponse, 1)
	stockAdvice := make(chan string, 1)
	recommendations := ticker.Recommendations

	go func() {
		company, _ := financial.Service().GetCompanyData(ticker.ID)
		advice, _ := gemini.Service().GetInvestmentAdvice(ticker.ID, company, recommendations)

		companyData <- company
		stockAdvice <- advice
	}()

	return stockItem{
		Ticker:          ticker.ID,
		Company:         <-companyData,
		Recommendations: FormatArrayWith(RecommendationResource{}, recommendations),
		Advice:          <-stockAdvice,
	}
}
