package main

import (
	"log"

	"github.com/MJU-Capstone-6/devmark-backend/internal/app"
)

func main() {
	application, err := app.InitApplication()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(application.App.Port)
	application.Run()
}
