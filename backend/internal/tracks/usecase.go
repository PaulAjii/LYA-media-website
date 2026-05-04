package tracks

import (
	"context"

	"github.com/PaulAjii/LYA-media-website/internal/storage"
	"github.com/PaulAjii/LYA-media-website/internal/tracks/dtos"
	"github.com/google/uuid"
)

type TrackUseCase struct {
	repo    *TrackRepository
	storage *storage.R2Storage
}

func NewUseCase(repo *TrackRepository, storage *storage.R2Storage) *TrackUseCase {
	return &TrackUseCase{
		repo:    repo,
		storage: storage,
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

func (u *TrackUseCase) DeleteTrack(ctx context.Context, trackID string) error {
	track, err := u.repo.GetByID(ctx, trackID)
	if err != nil {
		return err
	}

	// Delete file from R2 Storage
	err = u.storage.Delete(ctx, track.AudioURL)
	if err != nil {
		return err
	}
	return u.repo.Delete(ctx, trackID)
}
