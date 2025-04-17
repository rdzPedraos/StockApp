package models

type Ticker struct {
	ID      string `json:"id" gorm:"primaryKey;type:varchar(4)"`
	Company string `json:"company" gorm:"type:varchar(150);unique"`

	Recommendations []Recommendation `json:"recommendations"`
}
