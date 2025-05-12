package repository

import (
	"server/models"

	"gorm.io/gorm"
)

type GormBookRepository struct {
	db *gorm.DB
}

func NewGormBookRepository(db *gorm.DB) *GormBookRepository {
	return &GormBookRepository{db: db}
}

func (r *GormBookRepository) Create(book *models.Book) error {
	return r.db.Create(book).Error
}

func (r *GormBookRepository) FindByID(id uint) (*models.Book, error) {
	var book models.Book
	err := r.db.First(&book, id).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *GormBookRepository) FindAll() ([]*models.Book, error) {
	var books []*models.Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *GormBookRepository) Update(book *models.Book) error {
	return r.db.Save(book).Error
}

func (r *GormBookRepository) Delete(id uint) error {
	return r.db.Delete(&models.Book{}, id).Error
}
