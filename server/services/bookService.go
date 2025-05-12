package services

import (
	"errors"
	"server/database"
	"server/models"
	repository "server/repositories"
)

type BookService struct {
	repo models.BookRepository
}

func NewBookService() *BookService {
	bookRepository := repository.NewGormBookRepository(database.DB.Db)

	return &BookService{repo: bookRepository}
}

func (s *BookService) CreateBook(book *models.Book) error {
	if book.Status == "" {
		book.Status = models.BookStatusAvailable
	}

	return s.repo.Create(book)
}

func (s *BookService) GetBook(id uint) (*models.Book, error) {
	return s.repo.FindByID(id)
}

func (s *BookService) ListBooks() ([]*models.Book, error) {
	return s.repo.FindAll()
}

func (s *BookService) UpdateBook(book *models.Book) error {
	return s.repo.Update(book)
}

func (s *BookService) DeleteBook(id uint) error {
	existingBook, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if existingBook == nil {
		return errors.New("book not found")
	}

	return s.repo.Delete(id)
}

func (s *BookService) CheckBookStatus(status string) models.BookStatus {
	switch models.BookStatus(status) {
	case models.BookStatusAvailable, models.BookStatusBorrowed, models.BookStatusReserved, models.BookStatusLost:
		return models.BookStatus(status)
	default:
		return models.BookStatusAvailable
	}
}
