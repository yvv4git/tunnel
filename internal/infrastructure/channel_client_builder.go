package infrastructure

import (
	"context"
	"log/slog"

	"github.com/songgao/water"
	"github.com/yvv4git/tunnel/internal/infrastructure/config"
)

type ChannelClient interface {
	Start(ctx context.Context) error
	Close() error
}

type ChannelClientBuilder struct {
	logger    *slog.Logger
	cfg       config.Config
	tunDevice *water.Interface
}

func NewChannelClientBuilder(logger *slog.Logger, cfg config.Config, tunDevice *water.Interface) *ChannelClientBuilder {
	return &ChannelClientBuilder{
		logger:    logger,
		cfg:       cfg,
		tunDevice: tunDevice,
	}
}

func (b *ChannelClientBuilder) Build(channelType config.ChannelType) (ChannelClient, error) {
	switch channelType {
	case config.ChannelTCP:
		return NewClientTCP(b.logger, b.cfg.DirectConnection.Client, b.tunDevice), nil
	default:
		return nil, ErrInvalidChannelType
	}
}
