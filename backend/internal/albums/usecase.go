package albums

import (
	"context"

	"github.com/PaulAjii/LYA-media-website/internal/albums/dtos"
	"github.com/google/uuid"
)

type AlbumUseCase struct {
	repo *AlbumRepository
}

func NewUseCase(repo *AlbumRepository) *AlbumUseCase {
	return &AlbumUseCase{
		repo: repo,
	}
}

// Create Album
func (u *AlbumUseCase) Create(ctx context.Context, payload dtos.CreateAlbum) (*dtos.Album, error) {
	return u.repo.Create(ctx, payload)
}

// Get Albums
func (u *AlbumUseCase) GetAlbums(ctx context.Context) ([]dtos.Album, error) {
	return u.repo.GetAlbums(ctx)
}

// Get Album by ID
func (u *AlbumUseCase) GetAlbumByID(ctx context.Context, id uuid.UUID) (dtos.AlbumWithTracks, error) {
	album, err := u.repo.GetAlbumByID(ctx, id)
	if err != nil {
		return dtos.AlbumWithTracks{}, err
	}

	return album, nil
}

// Get Album by title
func (u *AlbumUseCase) GetAlbumByTitle(ctx context.Context, title string) ([]dtos.Album, error) {
	return u.repo.GetAlbumByTitle(ctx, title)
}

// Update Album by ID
func (u *AlbumUseCase) UpdateAlbumByID(ctx context.Context, id uuid.UUID, payload dtos.CreateAlbum) (*dtos.Album, error) {
	return u.repo.UpdateAlbumByID(ctx, id, payload)
}

// Delete Album by ID
func (u *AlbumUseCase) DeleteAlbumByID(ctx context.Context, id uuid.UUID) error {
	return u.repo.DeleteAlbumByID(ctx, id)
}
