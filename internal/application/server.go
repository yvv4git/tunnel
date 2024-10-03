package application

import (
	"context"

	"github.com/yvv4git/tunnel/internal/infrastructure"
)

type Server struct {
	application
	cfg infrastructure.Config
}

func NewServer(cfg infrastructure.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) start(ctx context.Context) error {
	s.log.Info("Starting ServerApplication")
	defer s.log.Info("Shutting down ServerApplication")

	// TODO: Implement server application logic

	return nil
}
