package tracks

import (
	"github.com/PaulAjii/LYA-media-website/internal/tracks/dtos"
	"github.com/PaulAjii/LYA-media-website/pkg/response"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type TrackHandler struct {
	u *TrackUseCase
}

func NewHandler(u *TrackUseCase) *TrackHandler {
	return &TrackHandler{
		u: u,
	}
}

func (h *TrackHandler) CreateTrack(c fiber.Ctx) error {
	var payload dtos.CreateTrackDTO
	if err := c.Bind().Body(&payload); err != nil {
		return response.Error(c, err.Error(), fiber.StatusBadRequest)
	}

	if payload.Title == "" || payload.AlbumID == uuid.Nil || payload.AudioURL == "" {
		return response.Error(c, "title, album ID, and audio URL are required", fiber.StatusBadRequest)
	}

	if payload.TrackNumber <= 0 {
		payload.TrackNumber = 1
	}

	track, err := h.u.Create(c.Context(), payload)
	if err != nil {
		return response.Error(c, err.Error(), fiber.StatusInternalServerError)
	}

	return response.Success(c, "Track created successfully", track, fiber.StatusCreated)
}

func (h *TrackHandler) GetTracksByAlbumID(c fiber.Ctx) error {
	albumID := c.Params("albumID")
	if albumID == "" {
		return response.Error(c, "AlbumID is required", fiber.StatusBadRequest)
	}

	id, err := uuid.Parse(albumID)
	if err != nil {
		return response.Error(c, "Invalid AlbumID format", fiber.StatusBadRequest)
	}

	tracks, err := h.u.GetTracks(c.Context(), id)
	if err != nil {
		return response.Error(c, err.Error(), fiber.StatusInternalServerError)
	}

	return response.Success(c, "Tracks retrieved successfully", tracks, fiber.StatusOK)
}
