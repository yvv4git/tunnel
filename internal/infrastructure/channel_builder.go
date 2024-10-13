package infrastructure

import (
	"context"

	"github.com/songgao/water"
)

type ChannelServer interface {
	Start(ctx context.Context) error
	Close() error
}

type ChannelType string

const (
	ChannelTCP ChannelType = "tcp"
	ChannelUDP ChannelType = "udp"
	// TODO: Add other types
)

type ChannelBuilder struct {
	cfg       Server
	tunDevice *water.Interface
}

func NewChannelBuilder(cfg Server, tunDevice *water.Interface) *ChannelBuilder {
	return &ChannelBuilder{
		cfg:       cfg,
		tunDevice: tunDevice,
	}
}

func (b *ChannelBuilder) Build(channelType ChannelType) (ChannelServer, error) {
	switch channelType {
	case ChannelTCP:
		return NewServerTCP(b.cfg, b.tunDevice), nil
	case ChannelUDP:
		return nil, nil // TODO: implement
	default:
		return nil, ErrInvalidChannelType
	}
}
