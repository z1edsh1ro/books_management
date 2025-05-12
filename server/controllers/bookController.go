package controllers

import (
	"fmt"
	"server/models"
	"server/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type BookHandler struct {
	service *services.BookService
}

func NewBookController(service *services.BookService) *BookHandler {
	return &BookHandler{service: service}
}

func (h *BookHandler) CreateBook(c *fiber.Ctx) error {
	file, err := c.FormFile("image")

	if err != nil {
		log.Warn("FILE ERROR")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "ERROR",
			"error":  err.Error(),
		})
	}

	fileName := file.Filename

	err = c.SaveFile(file, "./storages/"+fileName)

	if err != nil {
		log.Warn("CANNOT SAVE FILE\nERROR: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "ERROR",
			"error":  err.Error(),
		})
	}

	publishedAt := c.FormValue("published_at")
	convertPublishedAt, err := services.ConvertDate(publishedAt)

	// when cannot convert date format
	if err != nil {
		log.Warn("CANNOT CONVERT DATE")
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "Error",
			"message": err.Error(),
		})
	}

	title := c.FormValue("title")
	author := c.FormValue("author")
	description := c.FormValue("description")
	status := c.FormValue("status")
	bookStatus := h.service.CheckBookStatus(status)

	book := models.Book{
		Title:       title,
		Author:      author,
		Description: description,
		PublishedAt: convertPublishedAt,
		Status:      bookStatus,
		Image:       fileName,
	}

	if err := h.service.CreateBook(&book); err != nil {
		log.Warn("CANNOT CREATE BOOK")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "ERROR",
			"error":  err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "CREATED",
		"data":   book,
	})
}

func (h *BookHandler) GetBook(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "Error",
			"message": err.Error(),
		})
	}

	book, err := h.service.GetBook(uint(id))

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "Error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "SUCCESS",
		"data":   book,
	})
}

func (h *BookHandler) ListBooks(c *fiber.Ctx) error {
	books, err := h.service.ListBooks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "Error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "SUCCESS",
		"data":   books,
	})
}

func (h *BookHandler) UpdateBook(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		log.Warn("ID ERROR")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "Error",
			"message": err.Error(),
		})
	}

	var book models.Book

	if err := c.BodyParser(&book); err != nil {
		log.Warn("PAYLOAD ERROR, PAYLOAD:", book)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "Error",
			"message": err.Error(),
		})
	}

	bookUpdate, err := h.service.GetBook(uint(id))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "Error",
			"message": err.Error(),
		})
	}

	if bookUpdate == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "Error",
			"message": "Book not found",
		})
	}

	bookUpdate.Title = book.Title
	bookUpdate.Author = book.Author
	bookUpdate.Description = book.Description
	bookUpdate.PublishedAt = book.PublishedAt
	bookUpdate.Status = book.Status

	if err := h.service.UpdateBook(bookUpdate); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "Error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "UPDATED",
		"data":   bookUpdate,
	})
}

func (h *BookHandler) DeleteBook(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "ERROR",
			"message": err.Error(),
		})
	}

	if err := h.service.DeleteBook(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "ERROR",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "DELETED",
		"message": fmt.Sprintf("DELETED BOOK ID: %d", id),
	})
}
