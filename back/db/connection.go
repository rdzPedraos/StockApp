package db

import (
	"app/config"
	"app/utils"
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
	})

	return db
}

func initializeDB() *gorm.DB {
	var (
		cfg = config.Get()
		err error
	)

	conn, err := gorm.Open(
		postgres.Open(cfg.DB.GetDSN()),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	conn.AutoMigrate(utils.Schemas...)

	return conn
}
