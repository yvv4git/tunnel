package service

import (
	"context"

	"github.com/yvv4git/tunnel/internal/infrastructure"
)

type Client struct {
	clientTCP *infrastructure.ClientTCP
}

func NewClient(clientTCP *infrastructure.ClientTCP) *Client {
	return &Client{
		clientTCP: clientTCP,
	}
}

func (c *Client) Processing(ctx context.Context) error {
	return c.clientTCP.Start(ctx)
}
