package db

import (
	"github.com/bagashiz/simpler-bank/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open("host=localhost user=root password=password dbname=simpler_bank port=5432 sslmode=disable TimeZone=UTC"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	err = db.Debug().AutoMigrate(
		&models.User{},
		&models.Account{},
		&models.Transfer{},
		&models.Entry{},
		&models.Session{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetDB() *gorm.DB {
	return db
}
