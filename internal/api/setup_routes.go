package api

import (
	"github.com/PaulAjii/LYA-media-website/internal/albums"
	"github.com/PaulAjii/LYA-media-website/pkg/db"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("api/v1")

	pool, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// albums route setup
	albumsRepo := albums.NewRepository(pool)
	albumsUsecase := albums.NewUseCase(albumsRepo)
	albumsHandler := albums.NewHandler(albumsUsecase)
	albums.SetupRoutes(api, albumsHandler)
}
