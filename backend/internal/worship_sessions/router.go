package worshipsessions

import (
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(api fiber.Router, h *WorshipSessionHandler) {
	ws := api.Group("/worship-sessions")

	ws.Get("/", h.GetAllWorshipSessions)
	ws.Post("/", h.CreateWorshipSession)
	ws.Post("/backup-singers", h.CreateWorshipBackupSinger)
	ws.Post("/worship-songs", h.CreateWorshipSong)
	ws.Get("/:id", h.GetWorshipSessionByID)
}
