package service

import (
	"context"

	"github.com/yvv4git/tunnel/internal/infrastructure"
)

type Client struct {
	channelClient infrastructure.ChannelClient
}

func NewClient(channelClient infrastructure.ChannelClient) *Client {
	return &Client{
		channelClient: channelClient,
	}
}

func (c *Client) Processing(ctx context.Context) error {
	return c.channelClient.Start(ctx)
}
