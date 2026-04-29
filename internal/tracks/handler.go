package tracks

import (
	"strings"

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

func (h *TrackHandler) UpdateTrack(c fiber.Ctx) error {
	trackID := c.Params("trackID")
	if trackID == "" {
		return response.Error(c, "TrackID is required", fiber.StatusBadRequest)
	}

	id, err := uuid.Parse(trackID)
	if err != nil {
		return response.Error(c, "Invalid TrackID format", fiber.StatusBadRequest)
	}

	var payload dtos.UpdateTrackDTO
	if err := c.Bind().Body(&payload); err != nil {
		return response.Error(c, err.Error(), fiber.StatusBadRequest)
	}

	updatedTrack, err := h.u.UpdateTrack(c.Context(), id, payload)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return response.Error(c, "Track not found", fiber.StatusNotFound)
		}
		return response.Error(c, err.Error(), fiber.StatusInternalServerError)
	}

	return response.Success(c, "Tracked updated successfully", updatedTrack, fiber.StatusOK)
}

func (h *TrackHandler) DeleteTrack(c fiber.Ctx) error {
	trackID := c.Params("trackID")
	if trackID == "" {
		return response.Error(c, "TrackID is required", fiber.StatusBadRequest)
	}

	id, err := uuid.Parse(trackID)
	if err != nil {
		return response.Error(c, "Invalid TrackID format", fiber.StatusBadRequest)
	}

	err = h.u.DeleteTrack(c.Context(), id)
	if err != nil {
		return response.Error(c, err.Error(), fiber.StatusInternalServerError)
	}

	return response.Success(c, "Track deleted successfully", nil, fiber.StatusNoContent)
}
