package worshipsessions

import (
	"context"
	"fmt"

	"github.com/PaulAjii/LYA-media-website/internal/worship_sessions/dtos"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type WorshipSessionsRepository struct {
	db *pgxpool.Pool
}

func NewWorshipSessionsRepository(db *pgxpool.Pool) *WorshipSessionsRepository {
	return &WorshipSessionsRepository{
		db: db,
	}
}

func (r *WorshipSessionsRepository) CreateWorshipSession(ctx context.Context, payload dtos.WorshipSessionPayload) (*dtos.WorshipSessionWithoutAudio, error) {
	stmt := `
		INSERT INTO worship_sessions (date, worship_leader, audio_url, thumbnail_url)
		VALUES ($1, $2, $3, $4)
		RETURNING id, date, worship_leader, thumbnail_url, created_at
	`

	rows, err := r.db.Query(ctx, stmt, payload.Date, payload.WorshipLeader, payload.AudioURL, payload.ThumbnailURL)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	worshipSession, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[dtos.WorshipSessionWithoutAudio])
	if err != nil {
		return nil, fmt.Errorf("failed to create worship session: %w", err)
	}

	return &worshipSession, nil
}

func (r *WorshipSessionsRepository) GetAllWorshipSessions(ctx context.Context) ([]dtos.WorshipSessionDTO, error) {
	stmt := `
		SELECT * FROM worship_sessions
		ORDER BY date DESC
	`
	rows, err := r.db.Query(ctx, stmt)
	if err != nil {
		return []dtos.WorshipSessionDTO{}, fmt.Errorf("failed to execute query: %w", err)
	}

	worshipSessions, err := pgx.CollectRows(rows, pgx.RowToStructByName[dtos.WorshipSessionDTO])
	if err != nil {
		return []dtos.WorshipSessionDTO{}, fmt.Errorf("failed to fetch worship sessions: %w", err)
	}

	return worshipSessions, nil
}

func (r *WorshipSessionsRepository) GetWorshipSessionByID(ctx context.Context, id uuid.UUID) (*dtos.WorshipSessionFullDTO, error) {
	stmt := `
		SELECT ws.id, ws.date, ws.worship_leader, ws.audio_url, ws.thumbnail_url, ws.created_at,
    	COALESCE(
			(
				SELECT jsonb_agg(
					jsonb_build_object(
						'id', bs.id,
						'name', bs.name
					)
				)
				FROM worship_backup_singers bs
				WHERE bs.session_id = ws.id
			),
			'[]'::jsonb
		) AS worship_backup_singers,
		COALESCE(
			(
				SELECT jsonb_agg(
					jsonb_build_object(
						'id', w.id,
						'songTitle', w.song_title,
						'lyrics', w.lyrics,
						'songOrder', w.song_order,
						'createdAt', w.created_at
					) ORDER BY w.song_order ASC
				)
				FROM worship_songs w
				WHERE w.session_id = ws.id
			),
				'[]'::jsonb
			) AS worship_songs
		FROM worship_sessions ws
		WHERE ws.id = $1
	`

	rows, err := r.db.Query(ctx, stmt, id)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	worshipSession, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[dtos.WorshipSessionFullDTO])
	if err != nil {
		return nil, fmt.Errorf("failed to fetch worship session: %w", err)
	}

	return &worshipSession, nil
}

func (r *WorshipSessionsRepository) CreateBackupSinger(ctx context.Context, payload dtos.WorshipBackupSingerPayload) (*dtos.WorshipBackupSingerDTO, error) {
	stmt := `
		INSERT INTO worship_backup_singers(session_id, name)
		VALUES($1, $2)
		RETURNING id, name
	`

	rows, err := r.db.Query(ctx, stmt, payload.SessionID, payload.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	worshipBackupSinger, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[dtos.WorshipBackupSingerDTO])
	if err != nil {
		return nil, fmt.Errorf("failed to fetch the backup singer: %w", err)
	}

	return &worshipBackupSinger, nil
}

func (r *WorshipSessionsRepository) CreateWorshipSong(ctx context.Context, payload dtos.WorshipSongPayload) (*dtos.WorshipSongDTO, error) {
	stmt := `
		INSERT INTO worship_songs(session_id, song_title, lyrics, song_order)
		VALUES($1, $2, $3, $4)
		RETURNING id, song_title, lyrics, song_order, created_at
	`

	rows, err := r.db.Query(ctx, stmt, payload.SessionID, payload.SongTitle, payload.Lyrics, payload.SongOrder)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	worshipSong, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[dtos.WorshipSongDTO])
	if err != nil {
		return nil, fmt.Errorf("failed to fetch the backup singer: %w", err)
	}

	return &worshipSong, nil
}
