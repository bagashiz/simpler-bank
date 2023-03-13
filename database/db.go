package database

import (
	"github.com/bagashiz/simpler-bank/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB contains the database connection.
type DB struct {
	*gorm.DB
	error
}

// NewDB returns a new DB connection object.
func NewDB(dsn string) (*DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &DB{db, nil}, nil
}

// Migrate runs the database migrations.
func (db *DB) Migrate() error {
	err := db.AutoMigrate(
		//* commented out for now to avoid constraint errors
		// &models.User{},
		&models.Account{},
		&models.Transfer{},
		&models.Entry{},
	)

	return err
}
