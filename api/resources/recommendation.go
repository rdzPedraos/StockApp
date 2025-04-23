package resources

import (
	"app/models"
	"time"
)

type recommendationItem struct {
	Time          time.Time `json:"time"`
	Action        string    `json:"action"`
	RatingFrom    string    `json:"rating_from"`
	RatingFromTxt string    `json:"rating_from_txt"`
	RatingTo      string    `json:"rating_to"`
	RatingToTxt   string    `json:"rating_to_txt"`
	TargetFrom    float32   `json:"target_from"`
	TargetTo      float32   `json:"target_to"`
	Broker        string    `json:"broker"`
}

type RecommendationResource struct{}

func (r RecommendationResource) Format(recommendation models.Recommendation) interface{} {
	return recommendationItem{
		Time:          recommendation.Time,
		Action:        string(recommendation.Action),
		RatingFrom:    recommendation.RatingFrom.Label(),
		RatingFromTxt: string(recommendation.RatingFrom),
		RatingTo:      recommendation.RatingTo.Label(),
		RatingToTxt:   string(recommendation.RatingTo),
		TargetFrom:    recommendation.TargetFrom,
		TargetTo:      recommendation.TargetTo,
		Broker:        recommendation.Broker.Name,
	}
}
