package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/PaulAjii/LYA-media-website/internal/api"
	"github.com/PaulAjii/LYA-media-website/pkg/db"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	pool, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer pool.Close()
	log.Println("Connected to the database successfully")

	app := fiber.New(fiber.Config{
		AppName:   "LYA Media Website",
		BodyLimit: 50 * 1024 * 1024,
	})

	allowedOrigins := "*"

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{allowedOrigins},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowCredentials: false,
	}))

	app.Get("/api/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "API is healthy and running",
		})
	})

	api.SetupRoutes(app)

	serverErrors := make(chan error, 1)

	go func() {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
		log.Printf("BloansBooks API running on port %s", port)
		serverErrors <- app.Listen(":"+port, fiber.ListenConfig{
			EnablePrintRoutes: true,
		})
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		log.Printf("Error starting server: %v", err)
	case sig := <-quit:
		log.Printf("Received signal: %v. Shutting down server...", sig)
		if err := app.Shutdown(); err != nil {
			log.Printf("Error during server shutdown: %v", err)
		}
	}

	log.Println("Server gracefully stopped")
}
