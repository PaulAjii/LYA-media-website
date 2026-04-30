package dtos

import (
	"time"

	"github.com/google/uuid"
)

type ChoirMinistrationDTO struct {
	ID               uuid.UUID `json:"id" db:"id"`
	SongTitle        string    `json:"songTitle" db:"song_title"`
	MinistrationDate time.Time `json:"date" db:"date"`
	SongWriter       string    `json:"songWriter" db:"songwriter"`
	SongLyrics       *string   `json:"lyrics" db:"lyrics"`
	AudioURL         string    `json:"audioURL" db:"audio_url"`
	ThumbnailURL     *string   `json:"thumbnailURL" db:"thumbnail_url"`
	CreatedAt        time.Time `json:"createdAt" db:"created_at"`
}

type ChoirMinistrationPayload struct {
	SongTitle        string    `json:"songTitle" db:"song_title" validate:"required"`
	MinistrationDate time.Time `json:"date" db:"date" validate:"required"`
	SongWriter       string    `json:"songWriter" db:"songwriter" validate:"required"`
	SongLyrics       *string   `json:"lyrics" db:"lyrics" validate:"required"`
	AudioURL         string    `json:"audioURL" db:"audio_url" validate:"required,url"`
	ThumbnailURL     *string   `json:"thumbnailURL" db:"thumbnail_url" validate:"required,url"`
}
