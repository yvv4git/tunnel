package infrastructure

import (
	"context"
	"log/slog"

	"github.com/songgao/water"
	"github.com/yvv4git/tunnel/internal/infrastructure/config"
)

type ChannelServer interface {
	Start(ctx context.Context) error
	Close() error
}

type ChannelServerBuilder struct {
	logger    *slog.Logger
	cfg       config.Config
	tunDevice *water.Interface
}

func NewChannelServerBuilder(logger *slog.Logger, cfg config.Config, tunDevice *water.Interface) *ChannelServerBuilder {
	return &ChannelServerBuilder{
		logger:    logger,
		cfg:       cfg,
		tunDevice: tunDevice,
	}
}

func (b *ChannelServerBuilder) Build(channelType config.ChannelType) (ChannelServer, error) {
	switch channelType {
	case config.ChannelTCP:
		return NewServerTCP(b.logger, b.cfg.DirectConnection.Server, b.tunDevice), nil
	case config.ChannelUDP:
		return NewServerUDP(b.cfg.DirectConnection.Server, b.tunDevice), nil
	default:
		return nil, ErrInvalidChannelType
	}
}
