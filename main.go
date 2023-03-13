package main

import (
	"log"

	"github.com/bagashiz/simpler-bank/configs"
	"github.com/bagashiz/simpler-bank/database"
	"github.com/bagashiz/simpler-bank/server"
)

func main() {
	config, err := configs.NewConfig("configs")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := database.NewDB(config.GetDSN())
	if err != nil {
		log.Fatal(err)
	}

	err = conn.Migrate()
	if err != nil {
		log.Fatal(err)
	}

	server, err := server.NewServer(config)
	if err != nil {
		log.Fatal(err)
	}

	err = server.Start(config.HTTP_SERVER_ADDRESS)
	if err != nil {
		log.Fatal(err)
	}
}
