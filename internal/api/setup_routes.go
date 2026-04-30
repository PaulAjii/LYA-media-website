package api

import (
	"github.com/PaulAjii/LYA-media-website/internal/albums"
	cm "github.com/PaulAjii/LYA-media-website/internal/choir_ministration"
	"github.com/PaulAjii/LYA-media-website/internal/storage"
	"github.com/PaulAjii/LYA-media-website/internal/tracks"
	"github.com/PaulAjii/LYA-media-website/pkg/db"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("api/v1")

	pool, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	r2Storage := storage.NewR2Storage()

	// albums route setup
	albumsRepo := albums.NewRepository(pool)
	albumsUsecase := albums.NewUseCase(albumsRepo)
	albumsHandler := albums.NewHandler(albumsUsecase)
	albums.SetupRoutes(api, albumsHandler)

	// tracks route setup
	tracksRepo := tracks.NewTrackRepository(pool)
	tracksUsecase := tracks.NewUseCase(tracksRepo, r2Storage)
	tracksHandler := tracks.NewHandler(tracksUsecase, r2Storage)
	tracks.SetupRoutes(api, tracksHandler)

	// choir ministration route setup
	cmRepo := cm.NewChoirMinistrationRepo(pool)
	cmUsecase := cm.NewChoirMinistrationUseCase(cmRepo)
	cmHandler := cm.NewChoirMinistrationHandler(cmUsecase, r2Storage)
	cm.SetupRoutes(api, cmHandler)
}
