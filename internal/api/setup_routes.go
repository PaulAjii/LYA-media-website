package api

import (
	"os"

	"github.com/PaulAjii/LYA-media-website/internal/albums"
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

	supabaseStorage, err := storage.NewSupabaseStorage(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_SERVICE_ROLE_KEY"), os.Getenv("SUPABASE_BUCKET_NAME"))
	if err != nil {
		panic(err)
	}

	// albums route setup
	albumsRepo := albums.NewRepository(pool)
	albumsUsecase := albums.NewUseCase(albumsRepo)
	albumsHandler := albums.NewHandler(albumsUsecase)
	albums.SetupRoutes(api, albumsHandler)

	// tracks route setup
	tracksRepo := tracks.NewTrackRepository(pool)
	tracksUsecase := tracks.NewUseCase(tracksRepo)
	tracksHandler := tracks.NewHandler(tracksUsecase, supabaseStorage)
	tracks.SetupRoutes(api, tracksHandler)
}
