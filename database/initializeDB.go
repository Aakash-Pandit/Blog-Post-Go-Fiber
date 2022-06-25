package database

import (
	"fmt"
	"log"
	"sync"

	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var once sync.Once

var database *gorm.DB

func GetDatabase() *gorm.DB {
	return database
}

func InitializeDB(dbconfig *config.ApplicationConfig) {
	once.Do(
		func() {
			var err error

			psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				dbconfig.PostgresHost,
				dbconfig.PostgresPort,
				dbconfig.PostgresUser,
				dbconfig.PostgresPassword,
				dbconfig.PostgresDbName,
			)

			database, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
			if err != nil {
				log.Fatal(err)
			}

			log.Println("Database Connection established Successfully:", database.Name())
		},
	)
}
