package choirministration

import (
	"context"

	"github.com/PaulAjii/LYA-media-website/internal/choir_ministration/dtos"
	"github.com/google/uuid"
)

type ChoirMinistrationUseCase struct {
	repo *ChoirMinistrationRepo
}

func NewChoirMinistrationUseCase(repo *ChoirMinistrationRepo) *ChoirMinistrationUseCase {
	return &ChoirMinistrationUseCase{
		repo: repo,
	}
}

func (u *ChoirMinistrationUseCase) Create(ctx context.Context, payload dtos.ChoirMinistrationPayload) (*dtos.ChoirMinistrationDTO, error) {
	return u.repo.Create(ctx, payload)
}

func (u *ChoirMinistrationUseCase) Get(ctx context.Context, id uuid.UUID) (*dtos.ChoirMinistrationDTO, error) {
	return u.repo.Get(ctx, id)
}

func (u *ChoirMinistrationUseCase) GetAll(ctx context.Context, limit, offset int) ([]dtos.ChoirMinistrationDTO, error) {
	return u.repo.GetAll(ctx, limit, offset)
}

func (u *ChoirMinistrationUseCase) Count(ctx context.Context) (int, error) {
	return u.repo.Count(ctx)
}
