package discovery

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/erciktiburak/Lumina-Mesh/internal/messaging"
)

const (
	DiscoverySubject = "mesh.discovery"
)

type NodeInfo struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Timestamp time.Time `json:"timestamp"`
}

type DiscoveryService struct {
	client *messaging.Client
	nodeID string
}

func NewDiscoveryService(client *messaging.Client, nodeID string) *DiscoveryService {
	return &DiscoveryService{
		client: client,
		nodeID: nodeID,
	}
}

func (d *DiscoveryService) Start(ctx context.Context) error {
	_, err := d.client.Subscribe(DiscoverySubject, func(msg *nats.Msg) {
		var info NodeInfo
		if err := json.Unmarshal(msg.Data, &info); err == nil {
			if info.ID != d.nodeID {
				log.Printf("Discovered node: %s (Type: %s)", info.ID, info.Type)
			}
		}
	})
	if err != nil {
		return err
	}

	go d.announceLoop(ctx)
	return nil
}

func (d *DiscoveryService) announceLoop(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			info := NodeInfo{
				ID:        d.nodeID,
				Type:      "worker",
				Timestamp: time.Now(),
			}
			data, _ := json.Marshal(info)
			_ = d.client.Publish(DiscoverySubject, data)
		}
	}
}
