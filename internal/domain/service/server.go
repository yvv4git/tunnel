package service

import (
	"context"

	"github.com/yvv4git/tunnel/internal/infrastructure"
)

type Server struct {
	serverTCP *infrastructure.ServerTCP
}

func NewServer(serverTCP *infrastructure.ServerTCP) *Server {
	return &Server{
		serverTCP: serverTCP,
	}
}

func (s *Server) Processing(ctx context.Context) error {
	return s.serverTCP.Start(ctx)
}
