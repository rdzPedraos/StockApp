package models

import (
	"gorm.io/gorm"
)

type Ticker struct {
	ID      string `json:"id" gorm:"primaryKey;type:varchar(5)"`
	Company string `json:"company" gorm:"type:varchar(150)"`

	Recommendations []Recommendation `json:"recommendations"`
}

func (Ticker) SearchScope(search string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if search != "" {
			db = db.Where("id LIKE ? OR company LIKE ?", "%"+search+"%", "%"+search+"%")
		}

		return db
	}
}
