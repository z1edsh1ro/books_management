package main

import (
	"server/controllers"
	"server/services"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	bookService := services.NewBookService()
	bookController := controllers.NewBookController(bookService)

	// Book routes
	books := app.Group("/book")
	books.Post("/", bookController.CreateBook)
	books.Get("/", bookController.ListBooks)
	books.Get("/:id", bookController.GetBook)
	books.Put("/:id", bookController.UpdateBook)
	books.Delete("/:id", bookController.DeleteBook)

	app.Static("/files", "./storages")
}
