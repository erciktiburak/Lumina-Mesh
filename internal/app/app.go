package app

import (
	"log"

	"github.com/erciktiburak/Lumina-Mesh/internal/messaging"
)

type App struct {
	msgClient *messaging.Client
}

func New() *App {
	return &App{}
}

func (a *App) Run() error {
	log.Println("Lumina-Mesh Node is starting...")

	// Initialize messaging client (NATS)
	client, err := messaging.NewClient("")
	if err != nil {
		return err
	}
	a.msgClient = client
	defer a.msgClient.Close()

	log.Println("Lumina-Mesh Node is running and connected to messaging layer.")

	// Keep alive
	select {}
}
