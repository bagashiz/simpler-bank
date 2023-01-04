package main

import (
	"github.com/bagashiz/simpler-bank/configs"
	"github.com/bagashiz/simpler-bank/db"
	"github.com/bagashiz/simpler-bank/server"
)

func main() {
	// initiate database
	db.Connect()
	db.Migrate()

	// initiate server
	server.SetupRouter()
	server.Start(configs.GetHTTPServerAddress())
}
