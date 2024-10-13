package application

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/yvv4git/tunnel/internal/domain/service"
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
	c.log.Info("Starting client application")
	defer c.log.Info("Shutting down client application")

	tunDeviceBuilder, err := infrastructure.NewDeviceTUNClientBuilder(c.cfg)
	if err != nil {
		return fmt.Errorf("create TUN device: %w", err)
	}

	tunDevice, err := tunDeviceBuilder.Build(infrastructure.Platform(c.cfg.Server.DeviceTUN.Platform))
	if err != nil {
		return fmt.Errorf("build TUN device: %w", err)
	}
	defer tunDevice.Close()

	channelClientBuilder := infrastructure.NewChannelBuilder(c.cfg, tunDevice)
	channelClient, err := channelClientBuilder.BuildClient(c.cfg.Server.ChannelType)
	if err != nil {
		return fmt.Errorf("build client: %w", err)
	}
	defer channelClient.Close()

	svc := service.NewClient(channelClient)
	if err := svc.Processing(ctx); err != nil {
		return fmt.Errorf("start client: %w", err)
	}

	return nil
}
