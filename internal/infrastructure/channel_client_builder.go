package infrastructure

import (
	"context"

	"github.com/songgao/water"
)

type ChannelClient interface {
	Start(ctx context.Context) error
	Close() error
}

type ChannelClientBuilder struct {
	cfg       Config
	tunDevice *water.Interface
}

func NewChannelClientBuilder(cfg Config, tunDevice *water.Interface) *ChannelClientBuilder {
	return &ChannelClientBuilder{
		cfg:       cfg,
		tunDevice: tunDevice,
	}
}

func (b *ChannelClientBuilder) Build(channelType ChannelType) (ChannelClient, error) {
	switch channelType {
	case ChannelTCP:
		return NewClientTCP(b.cfg.Client, b.tunDevice), nil
	case ChannelUDP:
		return nil, nil // TODO: implement
	default:
		return nil, ErrInvalidChannelType
	}
}
