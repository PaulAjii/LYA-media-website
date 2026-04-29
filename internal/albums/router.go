package albums

import (
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(api fiber.Router, h *AlbumHandler) {
	albums := api.Group("/albums")

	albums.Get("/", h.GetAllAlbums)
	albums.Post("/", h.CreateAlbum)
	albums.Get("/search", h.SearchAlbumsByTitle)
	albums.Get("/:id", h.GetAlbumByID)
	albums.Patch("/:id", h.UpdateAlbum)
	albums.Delete("/:id", h.DeleteAlbum)
}
