package tracks

import (
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(api fiber.Router, h *TrackHandler) {
	tracks := api.Group("/tracks")

	tracks.Post("/", h.CreateTrack)
	tracks.Get("/album/:albumID", h.GetTracksByAlbumID)
	tracks.Patch("/:trackID", h.UpdateTrack)
	tracks.Delete("/:trackID", h.DeleteTrack)
}
