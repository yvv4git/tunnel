package infrastructure

import (
	"context"
	"log/slog"

	"github.com/songgao/water"
)

type ChannelClient interface {
	Start(ctx context.Context) error
	Close() error
}

type ChannelClientBuilder struct {
	logger    *slog.Logger
	cfg       Config
	tunDevice *water.Interface
}

func NewChannelClientBuilder(logger *slog.Logger, cfg Config, tunDevice *water.Interface) *ChannelClientBuilder {
	return &ChannelClientBuilder{
		logger:    logger,
		cfg:       cfg,
		tunDevice: tunDevice,
	}
}

func (b *ChannelClientBuilder) Build(channelType ChannelType) (ChannelClient, error) {
	switch channelType {
	case ChannelTCP:
		return NewClientTCP(b.logger, b.cfg.Client, b.tunDevice), nil
	case ChannelUDP:
		return NewClientUDP(b.cfg.Client, b.tunDevice), nil
	default:
		return nil, ErrInvalidChannelType
	}
}
