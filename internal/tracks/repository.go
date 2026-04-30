package tracks

import (
	"context"
	"fmt"

	"github.com/PaulAjii/LYA-media-website/internal/tracks/dtos"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TrackRepository struct {
	db *pgxpool.Pool
}

func NewTrackRepository(db *pgxpool.Pool) *TrackRepository {
	return &TrackRepository{
		db: db,
	}
}

// Create Track
func (r *TrackRepository) Create(ctx context.Context, payload dtos.CreateTrackDTO) (*dtos.TrackDTO, error) {
	stmt := `
		INSERT INTO tracks(album_id, title, track_number, audio_url)
		VALUES($1, $2, $3, $4)
		RETURNING id, album_id, title, track_number, audio_url, created_at
	`

	rows, err := r.db.Query(ctx, stmt, payload.AlbumID, payload.Title, payload.TrackNumber, payload.AudioURL)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	track, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[dtos.TrackDTO])
	if err != nil {
		return nil, fmt.Errorf("failed to collect row: %w", err)
	}

	return &track, nil
}

func (r *TrackRepository) GetByID(ctx context.Context, id string) (*dtos.TrackDTO, error) {
	stmt := `
		SELECT id, album_id, title, track_number, audio_url, created_at
		FROM tracks
		WHERE id = $1
	`
	rows, err := r.db.Query(ctx, stmt, id)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	track, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[dtos.TrackDTO])
	if err != nil {
		return nil, fmt.Errorf("failed to collect row: %w", err)
	}

	return &track, nil
}

func (r *TrackRepository) GetByAlbumID(ctx context.Context, albumID uuid.UUID) ([]dtos.TrackDTO, error) {
	stmt := `
		SELECT id, album_id, title, track_number, audio_url, created_at
		FROM tracks
		WHERE album_id = $1
		ORDER BY track_number ASC
	`
	rows, err := r.db.Query(ctx, stmt, albumID)
	if err != nil {
		return []dtos.TrackDTO{}, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	tracks, err := pgx.CollectRows(rows, pgx.RowToStructByName[dtos.TrackDTO])
	if err != nil {
		return []dtos.TrackDTO{}, fmt.Errorf("failed to collect rows: %w", err)
	}

	return tracks, nil
}

func (r *TrackRepository) UpdateTrack(ctx context.Context, id uuid.UUID, payload dtos.UpdateTrackDTO) (*dtos.TrackDTO, error) {
	stmt := `
		UPDATE tracks
		SET 
			title = COALESCE($1, title),
			track_number = COALESCE($2, track_number)
		WHERE id = $3
		RETURNING id, album_id, title, track_number, audio_url, created_at
	`

	rows, err := r.db.Query(ctx, stmt, payload.Title, payload.TrackNumber, id)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	upadtedTrack, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[dtos.TrackDTO])
	if err != nil {
		return nil, fmt.Errorf("failed to collect row: %w", err)
	}

	return &upadtedTrack, nil
}

func (r *TrackRepository) Delete(ctx context.Context, id string) error {
	stmt := `
		DELETE FROM tracks
		WHERE id = $1
	`
	_, err := r.db.Exec(ctx, stmt, id)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}
