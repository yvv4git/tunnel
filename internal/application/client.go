package application

import (
	"context"
	"log/slog"

	"github.com/yvv4git/tunnel/internal/infrastructure"
)

type Client struct {
	application
	cfg infrastructure.Config
}

func NewClient(log *slog.Logger, cfg infrastructure.Config) *Client {
	c := &Client{
		application: application{
			log: log,
		},
		cfg: cfg,
	}

	c.app = c
	return c
}

func (c *Client) start(ctx context.Context) error {
	c.log.Info("Starting ClientApplication")
	defer c.log.Info("Shutting down ClientApplication")

	// TODO: Implement client application logic

	return nil
}
