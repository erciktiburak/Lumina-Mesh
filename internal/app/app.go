package app

import (
	"context"
	"log"

	"github.com/erciktiburak/Lumina-Mesh/internal/discovery"
	"github.com/erciktiburak/Lumina-Mesh/internal/messaging"
	"github.com/erciktiburak/Lumina-Mesh/internal/worker"
)

type App struct {
	msgClient *messaging.Client
	discovery *discovery.DiscoveryService
	pool      *worker.Pool
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

	// Initialize and start worker pool
	a.pool = worker.NewPool(5, 100)
	a.pool.Start(context.Background())
	defer a.pool.Stop()

	// Initialize discovery service
	a.discovery = discovery.NewDiscoveryService(a.msgClient, "node-1")
	if err := a.discovery.Start(context.Background()); err != nil {
		return err
	}

	log.Println("Lumina-Mesh Node is running and connected to messaging layer.")

	// Keep alive
	select {}
}
