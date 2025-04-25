package main

import (
	"log"
	"testovoe/cmd/server/app"
)

func main() {
	app := app.Initialize()
	defer app.Close()

	log.Printf("Starting server on port %s", app.Config.DB_PORT)
	if err := app.Router.Run(":" + app.Config.DB_PORT); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
