package main

import (
	"RESTSAMPLE/database"
	"RESTSAMPLE/router"
	"log"
)

func main() {
	database.Setup()
	engine := router.Router()
	err := engine.Run("127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
}
