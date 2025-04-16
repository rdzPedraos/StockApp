package models

import (
	"time"
)

type Recommendation struct {
	ID uint `json:"id"`

	TickerID string `json:"ticker_id"`
	Ticker   Ticker

	Action string `json:"action"`

	BrokerID uint `json:"broker_id"`
	Broker   Broker

	TargetFrom float32 `json:"target_from"`
	TargetTo   float32 `json:"target_to"`

	RatingFrom string `json:"rating_from"`
	RatingTo   string `json:"rating_to"`

	Time time.Time `json:"time" gorm:"not null"`
}
