package choirministration

import (
	"fmt"
	"strings"
	"time"

	"github.com/PaulAjii/LYA-media-website/internal/choir_ministration/dtos"
	"github.com/PaulAjii/LYA-media-website/internal/storage"
	"github.com/PaulAjii/LYA-media-website/pkg/response"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type DataWitPagination struct {
	Data       interface{} `json:"pagedData"`
	Count      int         `json:"count"`
	TotalCount int         `json:"totalCount"`
	Limit      int         `json:"limit"`
	Offset     int         `json:"offset"`
}

type ChoirMinistrationHandler struct {
	u       *ChoirMinistrationUseCase
	storage *storage.R2Storage
}

func NewChoirMinistrationHandler(u *ChoirMinistrationUseCase, storage *storage.R2Storage) *ChoirMinistrationHandler {
	return &ChoirMinistrationHandler{
		u:       u,
		storage: storage,
	}
}

func (h *ChoirMinistrationHandler) Create(c fiber.Ctx) error {
	songTitle := c.FormValue("songTitle")
	if songTitle == "" {
		return response.Error(c, "song title is a required field", fiber.StatusBadRequest)
	}

	ministrationDate := c.FormValue("date")
	if ministrationDate == "" {
		return response.Error(c, "ministration date is a required field", fiber.StatusBadRequest)
	}

	date, err := time.Parse(time.RFC3339, ministrationDate)
	if err != nil {
		return err
	}

	songWriter := c.FormValue("songWriter")
	if songWriter == "" {
		return response.Error(c, "song writer is a required field", fiber.StatusBadRequest)
	}

	songLyrics := c.FormValue("lyrics")
	thumbnailURL := c.FormValue("thumbnailURL")

	fileTitle := fmt.Sprintf("%s - %s", songTitle, ministrationDate)

	file, err := c.FormFile("audio")
	if err != nil {
		return response.Error(c, "audio file is required", fiber.StatusBadRequest)
	}

	openedFile, err := file.Open()
	if err != nil {
		return response.Error(c, "failed to open audio file", fiber.StatusBadRequest)
	}

	audioURL, err := h.storage.UploadFile(c.Context(), openedFile, file, fmt.Sprintf("choir_ministrations/%s", ministrationDate), fileTitle)
	if err != err {
		return response.Error(c, fmt.Sprintf("failed to upload audio: %v", err), fiber.StatusInternalServerError)
	}

	payload := dtos.ChoirMinistrationPayload{
		SongTitle:        songTitle,
		MinistrationDate: date,
		SongWriter:       songWriter,
		SongLyrics:       &songLyrics,
		AudioURL:         audioURL,
		ThumbnailURL:     &thumbnailURL,
	}

	choirMinistration, err := h.u.Create(c.Context(), payload)
	if err != nil {
		return response.Error(c, err.Error(), fiber.StatusInternalServerError)
	}

	return response.Success(c, "Choir Ministration created successfully", choirMinistration, fiber.StatusCreated)
}

func (h *ChoirMinistrationHandler) Get(c fiber.Ctx) error {
	ministrationID := c.Params("id")
	if ministrationID == "" {
		return response.Error(c, "Choir Ministration ID is required", fiber.StatusBadRequest)
	}

	id, err := uuid.Parse(ministrationID)
	if err != nil {
		return response.Error(c, "Invalid ministration ID", fiber.StatusBadRequest)
	}

	ministration, err := h.u.Get(c.Context(), id)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return response.Error(c, "choir ministration not found", fiber.StatusNotFound)
		}
		return response.Error(c, err.Error(), fiber.StatusInternalServerError)
	}

	return response.Success(c, "Choir ministration fetched successfully", ministration, fiber.StatusOK)
}

func (h *ChoirMinistrationHandler) GetAll(c fiber.Ctx) error {
	totalCount, err := h.u.Count(c.Context())
	if err != nil {
		return response.Error(c, err.Error(), fiber.StatusInternalServerError)
	}

	limit := fiber.Query(c, "limit", 10)
	offset := fiber.Query(c, "offset", 0)

	ministrationsList, err := h.u.GetAll(c.Context(), limit, offset)
	if err != nil {
		return response.Error(c, err.Error(), fiber.StatusInternalServerError)
	}

	data := &DataWitPagination{
		Data:       ministrationsList,
		Count:      len(ministrationsList),
		TotalCount: totalCount,
		Limit:      limit,
		Offset:     offset,
	}

	return response.Success(c, "Choir ministrations fetched successfully", data, fiber.StatusOK)
}
