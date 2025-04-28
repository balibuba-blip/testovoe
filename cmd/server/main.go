package main

import (
	"log"
	"testovoe/cmd/server/app"
	"testovoe/config"
)

func main() {
	app := app.Initialize()
	defer app.Close()

	// Получаем порт из конфига
	port := config.GetServerPort()
	log.Printf("Starting server on port %s", port)

	if err := app.Router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
