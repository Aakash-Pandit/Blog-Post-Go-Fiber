package storage

import (
	"log"

	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/models"
	"gorm.io/gorm"
)

func MigrateDatabase(db *gorm.DB) error {
	err := models.MigrateBlogs(db)
	if err != nil {
		log.Fatal("could not migrate db")
		return err
	}

	err = models.MigrateUsers(db)
	if err != nil {
		log.Fatal("could not migrate db")
		return err
	}

	return nil
}
