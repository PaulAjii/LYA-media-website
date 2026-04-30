package tracks

import (
	"fmt"
	"strings"

	"github.com/PaulAjii/LYA-media-website/internal/storage"
	"github.com/PaulAjii/LYA-media-website/internal/tracks/dtos"
	"github.com/PaulAjii/LYA-media-website/pkg/response"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type TrackHandler struct {
	u       *TrackUseCase
	storage *storage.R2Storage
}

func NewHandler(u *TrackUseCase, storage *storage.R2Storage) *TrackHandler {
	return &TrackHandler{
		u:       u,
		storage: storage,
	}
}

// CreateTrack handles multipart form upload with audio file
func (h *TrackHandler) CreateTrack(c fiber.Ctx) error {
	// Parse form data
	albumIDStr := c.FormValue("albumId")
	if albumIDStr == "" {
		return response.Error(c, "albumId is required", fiber.StatusBadRequest)
	}

	albumID, err := uuid.Parse(albumIDStr)
	if err != nil {
		return response.Error(c, "invalid albumId format", fiber.StatusBadRequest)
	}

	albumTitle := c.FormValue("albumTitle")
	if albumTitle == "" {
		return response.Error(c, "albumTitle is required", fiber.StatusBadRequest)
	}

	title := c.FormValue("title")
	if title == "" {
		return response.Error(c, "title is required", fiber.StatusBadRequest)
	}

	trackNumberStr := c.FormValue("trackNumber")
	trackNumber := 1
	if trackNumberStr != "" {
		trackNumber = parseInt(trackNumberStr, 1)
	}

	// Parse track title to create a clean filename
	trackTitle := fmt.Sprintf("%s - Track %d - %s", albumTitle, trackNumber, title)

	// Get uploaded file
	file, err := c.FormFile("audio")
	if err != nil {
		return response.Error(c, "audio file is required", fiber.StatusBadRequest)
	}

	// Open the file
	openedFile, err := file.Open()
	if err != nil {
		return response.Error(c, "failed to open audio file", fiber.StatusBadRequest)
	}
	defer openedFile.Close()

	// Upload to R2 Storage
	audioURL, err := h.storage.UploadFile(c.Context(), openedFile, file, fmt.Sprintf("teachings/%s", albumTitle), trackTitle)
	if err != nil {
		return response.Error(c, fmt.Sprintf("failed to upload audio: %v", err), fiber.StatusInternalServerError)
	}

	// Create track with the uploaded URL
	payload := dtos.CreateTrackDTO{
		AlbumID:     albumID,
		AlbumTitle:  albumTitle,
		Title:       title,
		TrackNumber: trackNumber,
		AudioURL:    audioURL,
	}

	track, err := h.u.Create(c.Context(), payload)
	if err != nil {
		return response.Error(c, err.Error(), fiber.StatusInternalServerError)
	}

	return response.Success(c, "Track created successfully", track, fiber.StatusCreated)
}

func parseInt(s string, defaultValue int) int {
	var result int
	fmt.Sscanf(s, "%d", &result)
	if result == 0 {
		return defaultValue
	}
	return result
}

// GetTracksByAlbumID retrieves all tracks for a given album ID
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

// UpdateTrack allows updating track details except for the audio file
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

// DeleteTrack deletes a track and its associated audio file from R2 Storage
func (h *TrackHandler) DeleteTrack(c fiber.Ctx) error {
	trackID := c.Params("trackID")
	if trackID == "" {
		return response.Error(c, "TrackID is required", fiber.StatusBadRequest)
	}

	_, err := uuid.Parse(trackID)
	if err != nil {
		return response.Error(c, "Invalid TrackID format", fiber.StatusBadRequest)
	}

	err = h.u.DeleteTrack(c.Context(), trackID)
	if err != nil {
		return response.Error(c, err.Error(), fiber.StatusInternalServerError)
	}

	return response.Success(c, "Track deleted successfully", nil, fiber.StatusNoContent)
}
