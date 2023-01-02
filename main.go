package main

import "github.com/bagashiz/simpler-bank/db"

func main() {
	_, err := db.InitDB()
	if err != nil {
		panic(err.Error())
	}
}
