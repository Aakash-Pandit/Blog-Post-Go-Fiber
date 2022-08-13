package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthToken struct {
	Token string `json:"token" validate:"required"`
}

type User struct {
	ID        uuid.UUID `json:"id" gorm:"primary key"`
	FirstName string    `json:"first_name" validate:"required"`
	LastName  string    `json:"last_name" validate:"required"`
	Email     string    `json:"email" validate:"required"`
	Username  string    `json:"username"`
	Created   time.Time `json:"created"`
	Modified  time.Time `json:"modified"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	user.Created = time.Now()
	user.Modified = time.Now()
	return nil
}

func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {
	user.Modified = time.Now()
	return nil
}

func MigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	return err
}
