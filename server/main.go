package main

import (
	"server/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// connect database
	database.ConnectDB()

	// Create Fiber app
	app := fiber.New()

	// set up middleware
	setupMiddleware(app)

	// Setup routes
	setupRoutes(app)

	// Start server
	app.Listen(":3000")
}
