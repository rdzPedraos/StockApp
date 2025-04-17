package models

type Broker struct {
	ID   uint   `gorm:"primarykey"`
	Name string `json:"name" gorm:"type:varchar(150);unique"`

	Recommendations []Recommendation `json:"recommendations"`
}
