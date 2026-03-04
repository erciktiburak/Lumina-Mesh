package main

import (
	"log"

	"github.com/erciktiburak/Lumina-Mesh/internal/app"
)

func main() {
	log.Println("Starting Lumina-Mesh...")
	application := app.New()
	if err := application.Run(); err != nil {
		log.Fatalf("Lumina-Mesh failed to run: %v", err)
	}
}
