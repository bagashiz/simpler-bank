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

func Connect() {
	db, err = gorm.Open(postgres.Open(configs.GetDSN()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to DB.")
}

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

func GetDB() *gorm.DB {
	return db
}
