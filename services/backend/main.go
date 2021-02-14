package main

import (
	"backend/server"
	"log"
)

func main() {
	app := server.NewApp()

	if err := app.Run("8091"); err != nil {
		log.Fatalf("%s", err.Error())
	}
}