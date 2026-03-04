package discovery

import (
	"context"
	"encoding/json"
	"log"
	"sync"
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
	client    *messaging.Client
	nodeID    string
	nodes     map[string]time.Time
	mu        sync.RWMutex
	onNodeUp  func(id string)
	onNodeDown func(id string)
}

func NewDiscoveryService(client *messaging.Client, nodeID string) *DiscoveryService {
	return &DiscoveryService{
		client: client,
		nodeID: nodeID,
		nodes:  make(map[string]time.Time),
	}
}

func (d *DiscoveryService) SetHandlers(up, down func(string)) {
	d.onNodeUp = up
	d.onNodeDown = down
}

func (d *DiscoveryService) Start(ctx context.Context) error {
	_, err := d.client.Subscribe(DiscoverySubject, func(msg *nats.Msg) {
		var info NodeInfo
		if err := json.Unmarshal(msg.Data, &info); err == nil {
			if info.ID != d.nodeID {
				d.updateNode(info.ID)
			}
		}
	})
	if err != nil {
		return err
	}

	go d.announceLoop(ctx)
	go d.cleanupLoop(ctx)
	return nil
}

func (d *DiscoveryService) updateNode(id string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if _, exists := d.nodes[id]; !exists {
		log.Printf("[Discovery] New node detected: %s", id)
		if d.onNodeUp != nil {
			d.onNodeUp(id)
		}
	}
	d.nodes[id] = time.Now()
}

func (d *DiscoveryService) cleanupLoop(ctx context.Context) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			d.mu.Lock()
			for id, lastSeen := range d.nodes {
				if time.Since(lastSeen) > 15*time.Second {
					log.Printf("[Discovery] Node offline: %s", id)
					delete(d.nodes, id)
					if d.onNodeDown != nil {
						d.onNodeDown(id)
					}
				}
			}
			d.mu.Unlock()
		}
	}
}

func (d *DiscoveryService) announceLoop(ctx context.Context) {
...
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
