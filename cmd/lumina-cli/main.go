package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/erciktiburak/Lumina-Mesh/internal/messaging"
	"github.com/erciktiburak/Lumina-Mesh/internal/discovery"
	"github.com/nats-io/nats.go"
)

func main() {
	fmt.Println("Lumina-CLI: Monitoring Mesh Nodes...")

	client, err := messaging.NewClient("")
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer client.Close()

	_, err = client.Subscribe(discovery.DiscoverySubject, func(msg *nats.Msg) {
		fmt.Printf("Node update received: %s
", string(msg.Data))
	})
	if err != nil {
		log.Fatalf("Failed to subscribe to discovery: %v", err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}
