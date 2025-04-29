package database

import (
	"server/models"

	"gorm.io/gorm"
)

// GormBookRepository implements the BookRepository interface using GORM
type GormBookRepository struct {
	db *gorm.DB
}

// NewGormBookRepository creates a new instance of GormBookRepository
func NewGormBookRepository(db *gorm.DB) *GormBookRepository {
	return &GormBookRepository{db: db}
}

// Create implements BookRepository.Create
func (r *GormBookRepository) Create(book *models.Book) error {
	return r.db.Create(book).Error
}

// FindByID implements BookRepository.FindByID
func (r *GormBookRepository) FindByID(id uint) (*models.Book, error) {
	var book models.Book
	err := r.db.First(&book, id).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

// FindAll implements BookRepository.FindAll
func (r *GormBookRepository) FindAll() ([]*models.Book, error) {
	var books []*models.Book
	err := r.db.Find(&books).Error
	return books, err
}

// Update implements BookRepository.Update
func (r *GormBookRepository) Update(book *models.Book) error {
	return r.db.Save(book).Error
}

// Delete implements BookRepository.Delete
func (r *GormBookRepository) Delete(id uint) error {
	return r.db.Delete(&models.Book{}, id).Error
}

// FindByISBN implements BookRepository.FindByISBN
func (r *GormBookRepository) FindByISBN(isbn string) (*models.Book, error) {
	var book models.Book
	err := r.db.Where("isbn = ?", isbn).First(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

// FindByStatus implements BookRepository.FindByStatus
func (r *GormBookRepository) FindByStatus(status models.BookStatus) ([]*models.Book, error) {
	var books []*models.Book
	err := r.db.Where("status = ?", status).Find(&books).Error
	return books, err
}
