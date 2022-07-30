package blogs

import (
	"time"

	"github.com/google/uuid"
)

type Blog struct {
	Id         uuid.UUID `json:"id"`
	PostedAt   time.Time `json:"posted_at" validate:"required"`
	Content    string    `json:"content" validate:"required"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

func (blog *Blog) BeforeCreate() error {
	blog.Id = uuid.New()
	return nil
}
