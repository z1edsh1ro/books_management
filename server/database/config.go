package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

// Config holds database configuration
type Config struct {
	DSN string
}

// NewConfig creates a new database configuration
func NewConfig(dsn string) *Config {
	return &Config{
		DSN: dsn,
	}
}

// ConnectDB establishes a database connection and returns the GORM DB instance
func ConnectDB(config *Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(config.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate the schema
	if err := db.AutoMigrate(&Book{}); err != nil {
		return nil, err
	}

	return db, nil
}

// InitializeDB initializes the database connection and returns the repository
func InitializeDB(config *Config) (*GormBookRepository, error) {
	db, err := ConnectDB(config)
	if err != nil {
		return nil, err
	}

	return NewGormBookRepository(db), nil
} 