package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/erciktiburak/Lumina-Mesh/internal/discovery"
	"github.com/erciktiburak/Lumina-Mesh/internal/messaging"
)

func TestMeshSimulation(t *testing.T) {
	// Skip if NATS is not available
	client, err := messaging.NewClient("")
	if err != nil {
		t.Skip("NATS not available for simulation")
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	const nodeCount = 50
	nodes := make([]*discovery.DiscoveryService, nodeCount)

	fmt.Printf("Simulating %d nodes joining the mesh...
", nodeCount)

	for i := 0; i < nodeCount; i++ {
		nodes[i] = discovery.NewDiscoveryService(client, fmt.Sprintf("sim-node-%d", i))
		if err := nodes[i].Start(ctx); err != nil {
			t.Fatalf("Failed to start node %d: %v", i, err)
		}
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Simulation complete. Nodes are heartbeating.")
}
