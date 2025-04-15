package models

type Ticker struct {
	ID   string `json:"id" gorm:"primaryKey;type:varchar(4)"`
	Name string `json:"name" gorm:"type:varchar(150)"`

	Recommendations []Recommendation `json:"recommendations"`
}
