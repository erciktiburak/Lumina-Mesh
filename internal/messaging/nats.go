package messaging

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

type Client struct {
	Conn *nats.Conn
}

func NewClient(url string) (*Client, error) {
	if url == "" {
		url = nats.DefaultURL
	}
	nc, err := nats.Connect(url, nats.RetryOnFailedConnect(true), nats.MaxReconnects(10), nats.ReconnectWait(time.Second))
	if err != nil {
		return nil, err
	}

	log.Printf("Connected to NATS server at %s", nc.ConnectedUrl())

	return &Client{Conn: nc}, nil
}

func (c *Client) Close() {
	if c.Conn != nil {
		c.Conn.Close()
		log.Println("NATS connection closed")
	}
}

func (c *Client) Publish(subject string, data []byte) error {
	return c.Conn.Publish(subject, data)
}

func (c *Client) Subscribe(subject string, handler nats.MsgHandler) (*nats.Subscription, error) {
	return c.Conn.Subscribe(subject, handler)
}
