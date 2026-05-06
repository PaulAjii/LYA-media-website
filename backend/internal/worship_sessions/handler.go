package worshipsessions

import (
	"fmt"

	"github.com/PaulAjii/LYA-media-website/internal/storage"
	"github.com/PaulAjii/LYA-media-website/internal/worship_sessions/dtos"
	"github.com/PaulAjii/LYA-media-website/pkg/response"
	"github.com/PaulAjii/LYA-media-website/pkg/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type WorshipSessionHandler struct {
	u  *WorshipSessionsUsecase
	r2 *storage.R2Storage
}

func NewWorshipSessionHandler(u *WorshipSessionsUsecase, r2 *storage.R2Storage) *WorshipSessionHandler {
	return &WorshipSessionHandler{
		u:  u,
		r2: r2,
	}
}

func (h *WorshipSessionHandler) CreateWorshipSession(c fiber.Ctx) error {
	sessionDate := c.FormValue("date")
	if sessionDate == "" {
		return response.Error(c, "Date is a required field", fiber.StatusBadRequest)
	}

	date, err := utils.FormatDate(sessionDate)
	if err != nil {
		return response.Error(c, err.Error(), fiber.StatusBadRequest)
	}

	worshipLeader := c.FormValue("worshipLeader")
	if worshipLeader == "" {
		return response.Error(c, "Worship Leader is a required field", fiber.StatusBadRequest)
	}

	thumbnailURL := c.FormValue("thumbnailURL")

	file, err := c.FormFile("audio")
	if err != nil {
		return response.Error(c, "Audio file is required", fiber.StatusBadRequest)
	}

	fileTitle := fmt.Sprintf("Worship Session - %s", date)

	openedFile, err := file.Open()
	if err != nil {
		return response.Error(c, "unable to open file", fiber.StatusBadRequest)
	}
	defer openedFile.Close()

	audioURL, err := h.r2.UploadFile(c.Context(), openedFile, file, "worship-sessions", fileTitle)
	if err != nil {
		return response.Error(c, fmt.Sprintf("failed to upload audio: %s", err), fiber.StatusInternalServerError)
	}

	payload := dtos.WorshipSessionPayload{
		Date:          date,
		WorshipLeader: worshipLeader,
		AudioURL:      audioURL,
		ThumbnailURL:  &thumbnailURL,
	}

	worshipSession, err := h.u.CreateWorshipSession(c.Context(), payload)
	if err != nil {
		return response.Error(c, err.Error(), fiber.StatusInternalServerError)
	}

	return response.Success(c, "Worship session created successfully", worshipSession, fiber.StatusCreated)
}

func (h *WorshipSessionHandler) GetAllWorshipSessions(c fiber.Ctx) error {
	worshipSessions, err := h.u.GetAllWorshipSession(c.Context())
	if err != nil {
		return response.Error(c, err.Error(), fiber.StatusInternalServerError)
	}

	return response.Success(c, "Worship Sessions fetched successfully", worshipSessions, fiber.StatusOK)
}

func (h *WorshipSessionHandler) CreateWorshipBackupSinger(c fiber.Ctx) error {
	var payload dtos.WorshipBackupSingerPayload
	if err := c.Bind().Body(&payload); err != nil {
		return response.Error(c, err.Error(), fiber.StatusBadRequest)
	}

	worshipBackupSinger, err := h.u.CreateBackupSinger(c.Context(), payload)
	if err != nil {
		return response.Error(c, err.Error(), fiber.StatusInternalServerError)
	}

	return response.Success(c, "Worship backup singer created successfully", worshipBackupSinger, fiber.StatusCreated)
}

func (h *WorshipSessionHandler) CreateWorshipSong(c fiber.Ctx) error {
	var payload dtos.WorshipSongPayload
	if err := c.Bind().Body(&payload); err != nil {
		return response.Error(c, err.Error(), fiber.StatusBadRequest)
	}

	worshipSong, err := h.u.CreateWorshipSong(c.Context(), payload)
	if err != nil {
		return response.Error(c, err.Error(), fiber.StatusInternalServerError)
	}

	return response.Success(c, "Worship song created successfully", worshipSong, fiber.StatusCreated)
}

func (h *WorshipSessionHandler) GetWorshipSessionByID(c fiber.Ctx) error {
	paramsID := c.Params("id")
	if paramsID == "" {
		return response.Error(c, "Worship Session ID is required", fiber.StatusBadRequest)
	}

	id, err := uuid.Parse(paramsID)
	if err != nil {
		return response.Error(c, "Invalid ID format", fiber.StatusBadRequest)
	}

	worshipSession, err := h.u.GetWorshipSessionByID(c.Context(), id)
	if err != nil {
		return response.Error(c, err.Error(), fiber.StatusInternalServerError)
	}

	return response.Success(c, "Worship Session fetched successfully", worshipSession, fiber.StatusOK)
}
