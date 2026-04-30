package dtos

import (
	"time"

	"github.com/google/uuid"
)

type TrackDTO struct {
	ID          uuid.UUID `json:"id" db:"id"`
	AlbumID     uuid.UUID `json:"albumId" db:"album_id"`
	Title       string    `json:"title" db:"title"`
	TrackNumber int       `json:"trackNumber" db:"track_number"`
	AudioURL    string    `json:"audioUrl" db:"audio_url"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
}

type CreateTrackDTO struct {
	AlbumID     uuid.UUID `json:"albumId" validate:"required"`
	AlbumTitle  string    `json:"albumTitle" validate:"required"`
	Title       string    `json:"title" validate:"required"`
	TrackNumber int       `json:"trackNumber" validate:"required"`
	AudioURL    string    `json:"audioUrl" validate:"required,url"`
}

// CreateTrackWithFileDTO is used for multipart form uploads
// type CreateTrackWithFileDTO struct {
// 	AlbumID     uuid.UUID `form:"albumId" validate:"required"`
// 	AlbumTitle  string    `form:"albumTitle" validate:"required"`
// 	Title       string    `form:"title" validate:"required"`
// 	TrackNumber int       `form:"trackNumber" validate:"required"`
// }

type UpdateTrackDTO struct {
	Title       *string `json:"title,omitempty"`
	TrackNumber *int    `json:"trackNumber,omitempty"`
}
