package models

import (
	"time"

	"gorm.io/gorm"
)

// Book represents the domain entity for a book
type Book struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// Domain attributes
	Title       string     `json:"title" gorm:"type:varchar(255);not null"`
	Author      string     `json:"author" gorm:"type:varchar(255);not null"`
	Description string     `json:"description" gorm:"type:text"`
	PublishedAt time.Time  `json:"published_at"`
	Status      BookStatus `json:"status" gorm:"type:varchar(20);not null;default:'available'"`
}

// BookStatus represents the possible states of a book
type BookStatus string

const (
	BookStatusAvailable BookStatus = "available"
	BookStatusBorrowed  BookStatus = "borrowed"
	BookStatusReserved  BookStatus = "reserved"
	BookStatusLost      BookStatus = "lost"
)

// BookRepository defines the interface for book persistence
type BookRepository interface {
	Create(book *Book) error
	FindByID(id uint) (*Book, error)
	FindAll() ([]*Book, error)
	Update(book *Book) error
	Delete(id uint) error
	FindByISBN(isbn string) (*Book, error)
	FindByStatus(status BookStatus) ([]*Book, error)
}
