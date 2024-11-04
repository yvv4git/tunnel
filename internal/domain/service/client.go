package service

import (
	"context"

	"github.com/yvv4git/tunnel/internal/infrastructure/direct"
)

type Client struct {
	channelClient direct.ChannelClient
}

func NewClient(channelClient direct.ChannelClient) *Client {
	return &Client{
		channelClient: channelClient,
	}
}

func (c *Client) Processing(ctx context.Context) error {
	return c.channelClient.Start(ctx)
}
