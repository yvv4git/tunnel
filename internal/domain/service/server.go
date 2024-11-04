package service

import (
	"context"

	"github.com/yvv4git/tunnel/internal/infrastructure/direct"
)

type Server struct {
	channelServer direct.ChannelServer
}

func NewServer(channelServer direct.ChannelServer) *Server {
	return &Server{
		channelServer: channelServer,
	}
}

func (s *Server) Processing(ctx context.Context) error {
	return s.channelServer.Start(ctx)
}
