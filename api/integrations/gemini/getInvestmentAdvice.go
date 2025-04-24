package gemini

import (
	"app/integrations/financial"
	"app/lib/cache"
	"app/models"
	"fmt"
	"time"
)

func (s *service) GetInvestmentAdvice(ticker string, companyData financial.CompanyDataResponse, recommendations []models.Recommendation) (string, error) {
	var advice string

	cache.Once(
		"Gemini.GetInvestmentAdvice."+ticker,
		time.Hour*1,
		&advice,
		func() (*string, error) {
			companyDataString := fmt.Sprintf("Company data: %v", companyData)
			recommendationsString := fmt.Sprintf("Recommendations: %v", recommendations)

			prompts := []string{
				"You are a expert financial advisor",
				"I have some company data and recommendations about the stock " + ticker,
				"Here is the company data: " + companyDataString,
				"Here are the recommendations: " + recommendationsString,
				"Please give me a company overview, analyst recommendation and consice advice, about the stock " + ticker + " based on the data and recommendations given, be more specific and concise",
				"Use only markdown, dont use html or any other tags",
			}

			response, err := s.Ask(prompts)

			if err != nil {
				return nil, err
			}

			return &response, nil
		},
	)

	return advice, nil
}
