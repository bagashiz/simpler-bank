package db

import (
	"log"

	"github.com/bagashiz/simpler-bank/configs"
	"github.com/bagashiz/simpler-bank/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

// Connect initializes database session based on the specified data source name and configurations using gorm.
func Connect() {
	db, err = gorm.Open(postgres.Open(configs.GetDSN()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to DB.")
}

// Migrate runs auto migration for the specified entities from models package using gorm.
func Migrate() {
	err = db.AutoMigrate(
		&models.User{},
		&models.Account{},
		&models.Transfer{},
		&models.Entry{},
		&models.Session{},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB migration completed.")
}

// GetDB returns current database instance.
func GetDB() *gorm.DB {
	return db
}
