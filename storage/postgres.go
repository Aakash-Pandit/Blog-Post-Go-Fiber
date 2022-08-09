package storage

import (
	"fmt"

	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/core"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func GetDatabase() *gorm.DB {
	return database
}

func NewConnection(config *core.AppConfig) (*gorm.DB, error) {
	var err error
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName, config.DBSSLMode,
	)

	database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return database, err
	}

	return database, nil
}
