package service

import "context"

type ChannelClient interface {
	Start(ctx context.Context) error
}

type SpeedtestClient struct {
	channelClient ChannelClient
}

func NewSpeedtestClient(channelClient ChannelClient) *SpeedtestClient {
	return &SpeedtestClient{
		channelClient: channelClient,
	}
}

func (s *SpeedtestClient) Processing(ctx context.Context) error {
	return s.channelClient.Start(ctx)
}
