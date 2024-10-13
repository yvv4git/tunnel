package service

import (
	"context"

	"github.com/yvv4git/tunnel/internal/infrastructure"
)

type Server struct {
	channelServer infrastructure.ChannelServer
}

func NewServer(channelServer infrastructure.ChannelServer) *Server {
	return &Server{
		channelServer: channelServer,
	}
}

func (s *Server) Processing(ctx context.Context) error {
	return s.channelServer.Start(ctx)
}
