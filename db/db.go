package db

import (
	"log"

	"github.com/bagashiz/simpler-bank/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Connect() {
	db, err = gorm.Open(postgres.Open("host=localhost user=root password=password dbname=simpler_bank port=5432 sslmode=disable TimeZone=UTC"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB!")
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
		panic("Cannot migrate DB!")
	}

	log.Println("DB migration completed.")
}

func GetDB() *gorm.DB {
	return db
}
