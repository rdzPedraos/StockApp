package models

type Ticker struct {
	ID      string `json:"id" gorm:"primaryKey;type:varchar(5)"`
	Company string `json:"company" gorm:"type:varchar(150)"`

	Recommendations []Recommendation `json:"recommendations"`
}
