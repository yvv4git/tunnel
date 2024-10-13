package infrastructure

import (
	"context"

	"github.com/songgao/water"
)

type ChannelServer interface {
	Start(ctx context.Context) error
	Close() error
}

type ChanneServerBuilder struct {
	cfg       Config
	tunDevice *water.Interface
}

func NewChannelServerBuilder(cfg Config, tunDevice *water.Interface) *ChanneServerBuilder {
	return &ChanneServerBuilder{
		cfg:       cfg,
		tunDevice: tunDevice,
	}
}

func (b *ChanneServerBuilder) Build(channelType ChannelType) (ChannelServer, error) {
	switch channelType {
	case ChannelTCP:
		return NewServerTCP(b.cfg.Server, b.tunDevice), nil
	case ChannelUDP:
		return NewServerUDP(b.cfg.Server, b.tunDevice), nil
	default:
		return nil, ErrInvalidChannelType
	}
}
