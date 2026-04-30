package albums

import (
	"context"
	"fmt"

	"github.com/PaulAjii/LYA-media-website/internal/albums/dtos"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AlbumRepository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *AlbumRepository {
	return &AlbumRepository{
		db: db,
	}
}

// Create Album
func (r *AlbumRepository) Create(ctx context.Context, payload dtos.CreateAlbum) (*dtos.Album, error) {
	stmt := `
		INSERT INTO albums(title, date, summary, thumbnail_url)
		VALUES($1, $2, $3, $4)
		RETURNING id, title, date, summary, thumbnail_url, created_at
	`

	rows, err := r.db.Query(ctx, stmt, payload.Title, payload.Date, payload.Summary, payload.ThumbnailURL)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	defer rows.Close()

	album, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[dtos.Album])
	if err != nil {
		return nil, fmt.Errorf("failed to collect album: %w", err)
	}

	return &album, nil
}

// Get Albums
func (r *AlbumRepository) GetAlbums(ctx context.Context) ([]dtos.Album, error) {
	stmt := `
		SELECT
			id, title, date, summary, thumbnail_url, created_at
		FROM albums
		ORDER BY created_at DESC
	`
	rows, err := r.db.Query(ctx, stmt)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	albums, err := pgx.CollectRows(rows, pgx.RowToStructByName[dtos.Album])
	if err != nil {
		return nil, fmt.Errorf("failed to collect albums: %w", err)
	}

	return albums, nil
}

// Get Album by ID
func (r *AlbumRepository) GetAlbumByID(ctx context.Context, id uuid.UUID) (dtos.AlbumWithTracks, error) {
	stmt := `
		SELECT
			a.id, a.title, a.date, a.summary, a.thumbnail_url, a.created_at,
		COALESCE(
			json_agg(
				json_build_object(
					'id'	, t.id,
					'title', t.title,
					'albumId', t.album_id,
					'trackNumber', t.track_number,
					'audioUrl', t.audio_url,
					'createdAt', t.created_at
				) ORDER BY t.track_number ASC
			) FILTER (WHERE t.id IS NOT NULL),
			 '[]'::json
		) AS tracks
		FROM albums a
		LEFT JOIN tracks t ON a.id = t.album_id
		WHERE a.id = $1
		GROUP BY a.id
	`
	rows, err := r.db.Query(ctx, stmt, id)
	if err != nil {
		return dtos.AlbumWithTracks{}, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	album, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[dtos.AlbumWithTracks])
	if err != nil {
		return dtos.AlbumWithTracks{}, fmt.Errorf("failed to collect album: %w", err)
	}

	return album, nil
}

// Get Album by Title
func (r *AlbumRepository) GetAlbumByTitle(ctx context.Context, title string) ([]dtos.Album, error) {
	stmt := `
		SELECT DISTINCT
			a.id, a.title, a.date, a.summary, a.thumbnail_url, a.created_at
		FROM albums a
		LEFT JOIN tracks t ON a.id = t.album_id
		WHERE a.title ILIKE $1 OR t.title ILIKE $2 ORDER BY a.date DESC
	`
	rows, err := r.db.Query(ctx, stmt, "%"+title+"%", "%"+title+"%")
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	albums, err := pgx.CollectRows(rows, pgx.RowToStructByName[dtos.Album])
	if err != nil {
		return nil, fmt.Errorf("failed to collect albums: %w", err)
	}

	return albums, nil
}

// Update Album by ID
func (r *AlbumRepository) UpdateAlbumByID(ctx context.Context, id uuid.UUID, payload dtos.CreateAlbum) (*dtos.Album, error) {
	stmt := `
		UPDATE albums
		SET 
			title = COALESCE($1, title),
			date = COALESCE($2, date),
			summary = COALESCE($3, summary),
			thumbnail_url = COALESCE($4, thumbnail_url)
		WHERE id = $5
		RETURNING id, title, date, summary, thumbnail_url, created_at
	`
	rows, err := r.db.Query(ctx, stmt, payload.Title, payload.Date, payload.Summary, payload.ThumbnailURL, id)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	album, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[dtos.Album])
	if err != nil {
		return nil, fmt.Errorf("failed to collect album: %w", err)
	}

	return &album, nil
}

// Delete Album by ID
func (r *AlbumRepository) DeleteAlbumByID(ctx context.Context, id uuid.UUID) error {
	stmt := `
		DELETE FROM albums
		WHERE id = $1
	`
	_, err := r.db.Exec(ctx, stmt, id)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}
