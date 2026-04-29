package albums

import (
	"strings"

	"github.com/PaulAjii/LYA-media-website/internal/albums/dtos"
	"github.com/PaulAjii/LYA-media-website/pkg/response"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type AlbumHandler struct {
	u *AlbumUseCase
}

func NewHandler(u *AlbumUseCase) *AlbumHandler {
	return &AlbumHandler{
		u: u,
	}
}

func (h *AlbumHandler) CreateAlbum(c fiber.Ctx) error {
	var payload dtos.CreateAlbum
	if err := c.Bind().Body(&payload); err != nil {
		return response.Error(c, err.Error(), fiber.StatusBadRequest)
	}

	if payload.Title == "" {
		return response.Error(c, "Title is a required field", fiber.StatusBadRequest)
	}

	album, err := h.u.Create(c.Context(), payload)
	if err != nil {
		return response.Error(c, err.Error(), fiber.StatusInternalServerError)
	}

	return response.Success(c, "Album created succesfully", album, fiber.StatusCreated)
}

func (h *AlbumHandler) GetAllAlbums(c fiber.Ctx) error {
	albums, err := h.u.GetAlbums(c.Context())
	if err != nil {
		return response.Error(c, err.Error(), fiber.StatusInternalServerError)
	}

	return response.Success(c, "Albums fetched successfully", albums, fiber.StatusOK)
}

func (h *AlbumHandler) SearchAlbumsByTitle(c fiber.Ctx) error {
	title := c.Query("title")
	if title == "" {
		return response.Error(c, "Title query parameter is required", fiber.StatusBadRequest)
	}

	albums, err := h.u.GetAlbumByTitle(c.Context(), title)
	if err != nil {
		return response.Error(c, err.Error(), fiber.StatusInternalServerError)
	}

	return response.Success(c, "Albums fetched successfully", albums, fiber.StatusOK)
}

func (h *AlbumHandler) GetAlbumByID(c fiber.Ctx) error {
	paramID := c.Params("id")
	if paramID == "" {
		return response.Error(c, "AlbumID is required", fiber.StatusBadRequest)
	}

	id, err := uuid.Parse(paramID)
	if err != nil {
		return response.Error(c, "Invalid AlbumID format", fiber.StatusBadRequest)
	}

	album, err := h.u.GetAlbumByID(c.Context(), id)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return response.Error(c, "Album not found", fiber.StatusNotFound)
		}
		return response.Error(c, err.Error(), fiber.StatusInternalServerError)
	}

	return response.Success(c, "Album fetched successfully", album, fiber.StatusOK)
}

func (h *AlbumHandler) UpdateAlbum(c fiber.Ctx) error {
	paramID := c.Params("id")
	if paramID == "" {
		return response.Error(c, "AlbumID is required", fiber.StatusBadRequest)
	}

	id, err := uuid.Parse(paramID)
	if err != nil {
		return response.Error(c, "Invalid AlbumID format", fiber.StatusBadRequest)
	}

	var payload dtos.CreateAlbum
	if err := c.Bind().Body(&payload); err != nil {
		return response.Error(c, err.Error(), fiber.StatusBadRequest)
	}

	album, err := h.u.UpdateAlbumByID(c.Context(), id, payload)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return response.Error(c, "Album not found", fiber.StatusNotFound)
		}
		return response.Error(c, err.Error(), fiber.StatusInternalServerError)
	}

	return response.Success(c, "Album updated successfully", album, fiber.StatusOK)
}

func (h *AlbumHandler) DeleteAlbum(c fiber.Ctx) error {
	paramID := c.Params("id")
	if paramID == "" {
		return response.Error(c, "AlbumID is required", fiber.StatusBadRequest)
	}

	id, err := uuid.Parse(paramID)
	if err != nil {
		return response.Error(c, "Invalid AlbumID format", fiber.StatusBadRequest)
	}

	err = h.u.DeleteAlbumByID(c.Context(), id)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return response.Error(c, "Album not found", fiber.StatusNotFound)
		}
		return response.Error(c, err.Error(), fiber.StatusInternalServerError)
	}

	return nil
}
