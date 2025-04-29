package main

import (
	"server/database"
	"server/handlers"
	"server/services"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App, db *database.GormBookRepository) {
	// Create service and handler instances
	bookService := services.NewBookService(db)
	bookHandler := handlers.NewBookHandler(bookService)

	// Book routes
	books := app.Group("/books")
	books.Post("/", bookHandler.CreateBook)
	books.Get("/", bookHandler.ListBooks)
	books.Get("/:id", bookHandler.GetBook)
	books.Put("/:id", bookHandler.UpdateBook)
	books.Delete("/:id", bookHandler.DeleteBook)
	books.Post("/:id/borrow", bookHandler.BorrowBook)
	books.Post("/:id/return", bookHandler.ReturnBook)
}
