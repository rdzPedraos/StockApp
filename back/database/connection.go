package db

import (
	"app/config"
	"app/models"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func Get() *gorm.DB {
	once.Do(func() {
		db = initializeDB()
		AutoMigrate(db)
	})

	return db
}

func initializeDB() *gorm.DB {
	var err error

	conn, err := gorm.Open(
		postgres.Open(config.DataBase.GetDSN()),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatal("failed to connect database", err)
		//panic(err)
	}

	return conn
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(models.Schemas...)
}
