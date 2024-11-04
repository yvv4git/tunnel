package service

import "context"

type ChannelServer interface {
	Start(ctx context.Context) error
	Close() error
}

type SpeedtestServer struct {
	channelServer ChannelServer
}

func NewSpeedtestServer(channelServer ChannelServer) *SpeedtestServer {
	return &SpeedtestServer{
		channelServer: channelServer,
	}
}

func (s *SpeedtestServer) Processing(ctx context.Context) error {
	return s.channelServer.Start(ctx)
}
