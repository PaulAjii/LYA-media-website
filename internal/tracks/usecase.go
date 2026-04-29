package tracks

import (
	"context"

	"github.com/PaulAjii/LYA-media-website/internal/tracks/dtos"
	"github.com/google/uuid"
)

type TrackUseCase struct {
	repo *TrackRepository
}

func NewUseCase(repo *TrackRepository) *TrackUseCase {
	return &TrackUseCase{
		repo: repo,
	}
}

func (u *TrackUseCase) Create(ctx context.Context, payload dtos.CreateTrackDTO) (*dtos.TrackDTO, error) {
	return u.repo.Create(ctx, payload)
}

func (u *TrackUseCase) GetTracks(ctx context.Context, albumID uuid.UUID) ([]dtos.TrackDTO, error) {
	return u.repo.GetByAlbumID(ctx, albumID)
}

func (u *TrackUseCase) UpdateTrack(ctx context.Context, trackID uuid.UUID, payload dtos.UpdateTrackDTO) (*dtos.TrackDTO, error) {
	return u.repo.UpdateTrack(ctx, trackID, payload)
}

func (u *TrackUseCase) DeleteTrack(ctx context.Context, trackID uuid.UUID) error {
	return u.repo.DeleteTrack(ctx, trackID)
}
