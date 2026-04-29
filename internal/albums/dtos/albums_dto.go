package dtos

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/PaulAjii/LYA-media-website/internal/tracks/dtos"
	"github.com/google/uuid"
)

type TracksList []dtos.TrackDTO

func (t *TracksList) Scan(src any) error {
	bytes, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("expected []byte for tracks, got %T", src)
	}
	return json.Unmarshal(bytes, t)
}

type Album struct {
	ID           uuid.UUID `json:"id" db:"id"`
	Title        string    `json:"title" db:"title"`
	Date         time.Time `json:"date" db:"date"`
	Summary      *string   `json:"summary" db:"summary"`
	ThumbnailURL *string   `json:"thumbnail_url" db:"thumbnail_url"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
}

type AlbumWithTracksDTO struct {
	Album  `json:"album"`
	Tracks []dtos.TrackDTO `json:"tracks"`
}

type CreateAlbum struct {
	Title        string    `json:"title"`
	Date         time.Time `json:"date"`
	Summary      *string   `json:"summary"`
	ThumbnailURL *string   `json:"thumbnail_url"`
}

type AlbumWithTracks struct {
	ID           uuid.UUID  `json:"id" db:"id"`
	Title        string     `json:"title" db:"title"`
	Date         time.Time  `json:"date" db:"date"`
	Summary      *string    `json:"summary" db:"summary"`
	ThumbnailURL *string    `json:"thumbnail_url" db:"thumbnail_url"`
	CreatedAt    time.Time  `json:"createdAt" db:"created_at"`
	Tracks       TracksList `json:"tracks"`
}
