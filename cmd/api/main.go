package main

import (
	"log"
	"os"

	"github.com/PaulAjii/LYA-media-website/pkg/db"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	app := fiber.New()

	pool, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer pool.Close()
	log.Println("Connected to the database successfully")

	app.Get("/api/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "API is healthy and running",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen(":" + port))
}
