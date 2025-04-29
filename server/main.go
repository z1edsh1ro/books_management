package main

import (
	"log"
	"server/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize database
	dbConfig := database.NewConfig("books.db")
	bookRepo, err := database.InitializeDB(dbConfig)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Create Fiber app
	app := fiber.New()

	// Setup routes
	setupRoutes(app, bookRepo)

	// Start server
	app.Listen(":3000")
}
