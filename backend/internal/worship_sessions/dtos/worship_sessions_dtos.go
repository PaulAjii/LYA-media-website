package dtos

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type BackupSingersList []WorshipBackupSingerDTO
type WorshipSongsList []WorshipSongDTO

func (t *BackupSingersList) Scan(src any) error {
	bytes, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("expected []byte for backup singers, got %T", src)
	}
	return json.Unmarshal(bytes, t)
}

func (t *WorshipSongsList) Scan(src any) error {
	bytes, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("expected []byte for worship songs, got %T", src)
	}
	return json.Unmarshal(bytes, t)
}

type WorshipSessionDTO struct {
	ID            uuid.UUID `json:"id" db:"id"`
	Date          string    `json:"date" db:"date"`
	WorshipLeader string    `json:"worshipLeader" db:"worship_leader"`
	AudioURL      string    `json:"audioURL" db:"audio_url"`
	ThumbnailURL  *string   `json:"thumbnailURL" db:"thumbnail_url"`
	CreatedAt     time.Time `json:"createdAt" db:"created_at"`
}

type WorshipSessionWithoutAudio struct {
	ID            uuid.UUID `json:"id" db:"id"`
	Date          string    `json:"date" db:"date"`
	WorshipLeader string    `json:"worshipLeader" db:"worship_leader"`
	ThumbnailURL  *string   `json:"thumbnailURL" db:"thumbnail_url"`
	CreatedAt     time.Time `json:"createdAt" db:"created_at"`
}

type WorshipSessionPayload struct {
	Date          string  `json:"date"`
	WorshipLeader string  `json:"worshipLeader"`
	AudioURL      string  `json:"audioURL"`
	ThumbnailURL  *string `json:"thumbnailURL"`
}

type WorshipBackupSingerDTO struct {
	ID   uuid.UUID `json:"id" db:"id"`
	Name string    `json:"name" db:"name"`
}

type WorshipBackupSingerPayload struct {
	SessionID uuid.UUID `json:"sessionID" db:"session_id"`
	Name      string    `json:"name" db:"name"`
}

type WorshipSongDTO struct {
	ID        uuid.UUID `json:"id" db:"id"`
	SongTitle string    `json:"songTitle" db:"song_title"`
	Lyrics    *string   `json:"lyrics" db:"lyrics"`
	SongOrder int       `json:"songOrder" db:"song_order"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type WorshipSongPayload struct {
	SessionID uuid.UUID `json:"sessionID"`
	SongTitle string    `json:"songTitle"`
	Lyrics    *string   `json:"lyrics"`
	SongOrder int       `json:"songOrder"`
}

type WorshipSessionFullDTO struct {
	ID            uuid.UUID         `json:"id" db:"id"`
	Date          string            `json:"date" db:"date"`
	WorshipLeader string            `json:"worshipLeader" db:"worship_leader"`
	AudioURL      string            `json:"audioURL" db:"audio_url"`
	ThumbnailURL  *string           `json:"thumbnailURL" db:"thumbnail_url"`
	CreatedAt     time.Time         `json:"createdAt" db:"created_at"`
	BackupSingers BackupSingersList `json:"backupSingers" db:"worship_backup_singers"`
	WorshipSongs  WorshipSongsList  `json:"worshipSongs" db:"worship_songs"`
}
