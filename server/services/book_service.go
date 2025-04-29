package services

import (
	"errors"
	"server/models"
)

// BookService handles business logic for books
type BookService struct {
	repo models.BookRepository
}

// NewBookService creates a new instance of BookService
func NewBookService(repo models.BookRepository) *BookService {
	return &BookService{repo: repo}
}

// CreateBook handles the business logic for creating a new book
func (s *BookService) CreateBook(book *models.Book) error {
	// Set default status if not provided
	if book.Status == "" {
		book.Status = models.BookStatusAvailable
	}

	return s.repo.Create(book)
}

// GetBook retrieves a book by ID
func (s *BookService) GetBook(id uint) (*models.Book, error) {
	return s.repo.FindByID(id)
}

// ListBooks retrieves all books
func (s *BookService) ListBooks() ([]*models.Book, error) {
	return s.repo.FindAll()
}

// UpdateBook handles the business logic for updating a book
func (s *BookService) UpdateBook(book *models.Book) error {
	// Check if book exists
	existingBook, err := s.repo.FindByID(book.ID)
	if err != nil {
		return err
	}
	if existingBook == nil {
		return errors.New("book not found")
	}

	return s.repo.Update(book)
}

// DeleteBook handles the business logic for deleting a book
func (s *BookService) DeleteBook(id uint) error {
	// Check if book exists
	existingBook, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if existingBook == nil {
		return errors.New("book not found")
	}

	return s.repo.Delete(id)
}

// BorrowBook handles the business logic for borrowing a book
func (s *BookService) BorrowBook(id uint) error {
	book, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if book == nil {
		return errors.New("book not found")
	}
	if book.Status != models.BookStatusAvailable {
		return errors.New("book is not available for borrowing")
	}

	book.Status = models.BookStatusBorrowed
	return s.repo.Update(book)
}

// ReturnBook handles the business logic for returning a book
func (s *BookService) ReturnBook(id uint) error {
	book, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if book == nil {
		return errors.New("book not found")
	}
	if book.Status != models.BookStatusBorrowed {
		return errors.New("book is not currently borrowed")
	}

	book.Status = models.BookStatusAvailable
	return s.repo.Update(book)
}
