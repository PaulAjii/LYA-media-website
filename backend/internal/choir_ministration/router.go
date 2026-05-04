package choirministration

import "github.com/gofiber/fiber/v3"

func SetupRoutes(api fiber.Router, h *ChoirMinistrationHandler) {
	choirMinistration := api.Group("/choir-ministrations")

	choirMinistration.Post("/", h.Create)
	choirMinistration.Get("/", h.GetAll)
	choirMinistration.Get("/:id", h.Get)
}
