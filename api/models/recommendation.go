package models

import (
	"time"

	"app/enums/actions"
	"app/enums/ratings"
)

type Recommendation struct {
	ID uint `json:"id"`

	TickerID string `json:"ticker_id" gorm:"not null;uniqueIndex:idx_rec"`
	Ticker   Ticker

	Action actions.Action `json:"action" gorm:"not null"`

	BrokerID uint `json:"broker_id" gorm:"not null;uniqueIndex:idx_rec"`
	Broker   Broker

	TargetFrom float32 `json:"target_from"`
	TargetTo   float32 `json:"target_to"`

	RatingFrom ratings.Rating `json:"rating_from"`
	RatingTo   ratings.Rating `json:"rating_to"`

	Time time.Time `json:"time" gorm:"not null;uniqueIndex:idx_rec"`
}
