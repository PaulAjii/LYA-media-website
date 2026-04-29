package dtos

import (
	"time"

	"github.com/google/uuid"
)

type Album struct {
	ID           uuid.UUID `json:"id" db:"id"`
	Title        string    `json:"title" db:"title"`
	Date         time.Time `json:"date" db:"date"`
	Summary      *string   `json:"summary" db:"summary"`
	ThumbnailURL *string   `json:"thumbnail_url" db:"thumbnail_url"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
}

type CreateAlbum struct {
	Title        string    `json:"title"`
	Date         time.Time `json:"date"`
	Summary      *string   `json:"summary"`
	ThumbnailURL *string   `json:"thumbnail_url"`
}
