package application

import (
	"context"

	"github.com/yvv4git/tunnel/internal/infrastructure"
)

type Client struct {
	application
	cfg infrastructure.Config
}

func NewClient(cfg infrastructure.Config) *Client {
	return &Client{
		cfg: cfg,
	}
}

func (c *Client) start(ctx context.Context) error {
	c.log.Info("Starting ClientApplication")
	defer c.log.Info("Shutting down ClientApplication")

	// TODO: Implement client application logic

	return nil
}
