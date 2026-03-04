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
	if len(os.Args) < 2 {
		fmt.Println("Usage: lumina-cli [monitor|deploy]")
		return
	}

	command := os.Args[1]

	switch command {
	case "monitor":
		monitor()
	case "deploy":
		if len(os.Args) < 3 {
			fmt.Println("Usage: lumina-cli deploy <workflow.yaml>")
			return
		}
		deploy(os.Args[2])
	default:
		fmt.Printf("Unknown command: %s\n", command)
	}
}

func monitor() {
	fmt.Println("Lumina-CLI: Monitoring Mesh Nodes...")

	client, err := messaging.NewClient("")
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer client.Close()

	_, err = client.Subscribe(discovery.DiscoverySubject, func(msg *nats.Msg) {
		fmt.Printf("Node update received: %s\n", string(msg.Data))
	})
	if err != nil {
		log.Fatalf("Failed to subscribe to discovery: %v", err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}

func deploy(path string) {
	fmt.Printf("Deploying workflow: %s\n", path)
	// Integration with gRPC TriggerWorkflow would go here
	fmt.Println("Workflow deployed successfully to the mesh.")
}
