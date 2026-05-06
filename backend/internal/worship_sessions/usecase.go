package worshipsessions

import (
	"context"

	"github.com/PaulAjii/LYA-media-website/internal/worship_sessions/dtos"
	"github.com/google/uuid"
)

type WorshipSessionsUsecase struct {
	repo *WorshipSessionsRepository
}

func NewWorshipSessionUsecase(repo *WorshipSessionsRepository) *WorshipSessionsUsecase {
	return &WorshipSessionsUsecase{
		repo: repo,
	}
}

func (u *WorshipSessionsUsecase) CreateWorshipSession(ctx context.Context, payload dtos.WorshipSessionPayload) (*dtos.WorshipSessionWithoutAudio, error) {
	return u.repo.CreateWorshipSession(ctx, payload)
}

func (u *WorshipSessionsUsecase) GetAllWorshipSession(ctx context.Context) ([]dtos.WorshipSessionDTO, error) {
	return u.repo.GetAllWorshipSessions(ctx)
}

func (u *WorshipSessionsUsecase) GetWorshipSessionByID(ctx context.Context, id uuid.UUID) (*dtos.WorshipSessionFullDTO, error) {
	return u.repo.GetWorshipSessionByID(ctx, id)
}

func (u *WorshipSessionsUsecase) CreateBackupSinger(ctx context.Context, payload dtos.WorshipBackupSingerPayload) (*dtos.WorshipBackupSingerDTO, error) {
	return u.repo.CreateBackupSinger(ctx, payload)
}

func (u *WorshipSessionsUsecase) CreateWorshipSong(ctx context.Context, payload dtos.WorshipSongPayload) (*dtos.WorshipSongDTO, error) {
	return u.repo.CreateWorshipSong(ctx, payload)
}
