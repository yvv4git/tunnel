package infrastructure

import (
	"context"

	"github.com/songgao/water"
)

type ChannelServer interface {
	Start(ctx context.Context) error
	Close() error
}

type ChannelClient interface {
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
	cfg       Config
	tunDevice *water.Interface
}

func NewChannelBuilder(cfg Config, tunDevice *water.Interface) *ChannelBuilder {
	return &ChannelBuilder{
		cfg:       cfg,
		tunDevice: tunDevice,
	}
}

func (b *ChannelBuilder) BuildServer(channelType ChannelType) (ChannelServer, error) {
	switch channelType {
	case ChannelTCP:
		return NewServerTCP(b.cfg.Server, b.tunDevice), nil
	case ChannelUDP:
		return nil, nil // TODO: implement
	default:
		return nil, ErrInvalidChannelType
	}
}

func (b *ChannelBuilder) BuildClient(channelType ChannelType) (ChannelClient, error) {
	switch channelType {
	case ChannelTCP:
		return NewClientTCP(b.cfg.Client, b.tunDevice), nil
	case ChannelUDP:
		return nil, nil // TODO: implement
	default:
		return nil, ErrInvalidChannelType
	}
}
