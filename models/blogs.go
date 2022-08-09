package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Blog struct {
	ID       uuid.UUID `json:"id" gorm:"primary key"`
	Title    string    `json:"title" validate:"required"`
	Content  string    `json:"content" validate:"required"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

// func (blog *Blog) BeforeCreate() error {
// 	(*blog).ID = uuid.New()
// 	return nil
// }

func (blog *Blog) BeforeCreate(tx *gorm.DB) (err error) {
	blog.ID = uuid.New()

	return nil
}

func MigrateBlogs(db *gorm.DB) error {
	err := db.AutoMigrate(&Blog{})
	return err
}
