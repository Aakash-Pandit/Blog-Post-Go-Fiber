package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Blog struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" validate:"required"`
	Content     string    `json:"content" validate:"required"`
	CreatedById uuid.UUID `json:"created_by_id" gorm:"type:uuid;column:user_foreign_key;not null;"`
	Created     time.Time `json:"created"`
	Modified    time.Time `json:"modified"`
}

func (blog *Blog) BeforeCreate(tx *gorm.DB) (err error) {
	blog.ID = uuid.New()
	blog.Created = time.Now()
	blog.Modified = time.Now()
	return nil
}

func (blog *Blog) BeforeUpdate(tx *gorm.DB) (err error) {
	blog.Modified = time.Now()
	return nil
}

func MigrateBlogs(db *gorm.DB) error {
	err := db.AutoMigrate(&Blog{})
	return err
}
