package config

import (
	"app/lib/helper"
	"fmt"
	"os"
)

type dbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func (c *dbConfig) GetDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode)
}

var DataBase = dbConfig{
	Host:     os.Getenv("DB_HOST"),
	Port:     os.Getenv("DB_PORT"),
	User:     os.Getenv("DB_USER"),
	Password: os.Getenv("DB_PASSWORD"),
	DBName:   os.Getenv("DB_NAME"),
	SSLMode:  helper.Coalesce(os.Getenv("DB_SSLMODE"), "verify-full"),
}
