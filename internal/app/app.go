package app

import (
	"context"
	"log"

	"github.com/erciktiburak/Lumina-Mesh/internal/discovery"
	"github.com/erciktiburak/Lumina-Mesh/internal/messaging"
)

type App struct {
	msgClient *messaging.Client
	discovery *discovery.DiscoveryService
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

	// Initialize discovery service
	a.discovery = discovery.NewDiscoveryService(a.msgClient, "node-1")
	if err := a.discovery.Start(context.Background()); err != nil {
		return err
	}

	log.Println("Lumina-Mesh Node is running and connected to messaging layer.")

	// Keep alive
	select {}
}
