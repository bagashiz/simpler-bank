package main

import "github.com/bagashiz/simpler-bank/db"

func main() {
	// initiate database
	db.Connect()
	db.Migrate()
}
