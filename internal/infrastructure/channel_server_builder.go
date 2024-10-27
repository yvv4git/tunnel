package infrastructure

import (
	"context"
	"log/slog"

	"github.com/songgao/water"
)

type ChannelServer interface {
	Start(ctx context.Context) error
	Close() error
}

type ChannelServerBuilder struct {
	logger    *slog.Logger
	cfg       Config
	tunDevice *water.Interface
}

func NewChannelServerBuilder(logger *slog.Logger, cfg Config, tunDevice *water.Interface) *ChannelServerBuilder {
	return &ChannelServerBuilder{
		logger:    logger,
		cfg:       cfg,
		tunDevice: tunDevice,
	}
}

func (b *ChannelServerBuilder) Build(channelType ChannelType) (ChannelServer, error) {
	switch channelType {
	case ChannelTCP:
		return NewServerTCP(b.logger, b.cfg.Server, b.tunDevice), nil
	case ChannelUDP:
		return NewServerUDP(b.cfg.Server, b.tunDevice), nil
	default:
		return nil, ErrInvalidChannelType
	}
}
