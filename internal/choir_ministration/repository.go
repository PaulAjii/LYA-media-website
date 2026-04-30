package choirministration

import (
	"context"
	"fmt"

	"github.com/PaulAjii/LYA-media-website/internal/choir_ministration/dtos"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ChoirMinistrationRepo struct {
	db *pgxpool.Pool
}

func NewChoirMinistrationRepo(db *pgxpool.Pool) *ChoirMinistrationRepo {
	return &ChoirMinistrationRepo{
		db: db,
	}
}

func (r *ChoirMinistrationRepo) Create(ctx context.Context, payload dtos.ChoirMinistrationPayload) (*dtos.ChoirMinistrationDTO, error) {
	stmt := `
		INSERT INTO choir_ministrations(song_title, date, songwriter, lyrics, audio_url, thumbnail_url)
		VALUES (@song_title, @date, @songwriter, @lyrics, @audio_url, @thumbnail_url)
		RETURNING id, song_title, date, songwriter, lyrics, audio_url, thumbnail_url, created_at
	`

	rows, err := r.db.Query(ctx, stmt, pgx.NamedArgs{
		"song_title":    payload.SongTitle,
		"date":          payload.MinistrationDate,
		"songwriter":    payload.SongWriter,
		"lyrics":        &payload.SongLyrics,
		"audio_url":     payload.AudioURL,
		"thumbnail_url": &payload.ThumbnailURL,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	choirMinistration, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[dtos.ChoirMinistrationDTO])
	if err != nil {
		return nil, fmt.Errorf("failed to create choir ministration: %w", err)
	}

	return &choirMinistration, nil
}

func (r *ChoirMinistrationRepo) Get(ctx context.Context, id uuid.UUID) (*dtos.ChoirMinistrationDTO, error) {
	stmt := `
		SELECT id, song_title, date, songwriter, lyrics, audio_url, thumbnail_url, created_at
		FROM choir_ministrations
		WHERE id = $1
	`

	rows, err := r.db.Query(ctx, stmt, id)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	choirMinistration, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[dtos.ChoirMinistrationDTO])
	if err != nil {
		return nil, fmt.Errorf("failed to fetch choir ministration: %w", err)
	}

	return &choirMinistration, nil
}

func (r *ChoirMinistrationRepo) GetAll(ctx context.Context, limit, offset int) ([]dtos.ChoirMinistrationDTO, error) {
	stmt := `
		SELECT id, song_title, date, songwriter, lyrics, audio_url, thumbnail_url, created_at
		FROM choir_ministrations
		ORDER BY date DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, stmt, limit, offset)
	if err != nil {
		return []dtos.ChoirMinistrationDTO{}, fmt.Errorf("failed to execute query: %w", err)
	}

	choirMinistrations, err := pgx.CollectRows(rows, pgx.RowToStructByName[dtos.ChoirMinistrationDTO])
	if err != nil {
		return []dtos.ChoirMinistrationDTO{}, fmt.Errorf("failed to fetch choir ministrations: %w", err)
	}

	return choirMinistrations, nil
}

func (r *ChoirMinistrationRepo) Count(ctx context.Context) (int, error) {
	stmt := `SELECT COUNT(*) FROM choir_ministrations`

	var count int
	rows, err := r.db.Query(ctx, stmt)
	if err != nil {
		return 0, fmt.Errorf("failed to count choir ministrations: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return 0, fmt.Errorf("failed to scan count: %w", err)
		}
	}

	return count, nil
}
