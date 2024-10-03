package application

import (
	"context"
	"log/slog"

	"github.com/yvv4git/tunnel/internal/infrastructure"
)

type Server struct {
	application
	cfg infrastructure.Config
}

func NewServer(log *slog.Logger, cfg infrastructure.Config) *Server {
	s := &Server{
		application: application{
			log: log,
		},
		cfg: cfg,
	}

	s.app = s
	return s
}

func (s *Server) start(ctx context.Context) error {
	s.log.Info("Starting ServerApplication")
	defer s.log.Info("Shutting down ServerApplication")

	// TODO: Implement server application logic

	return nil
}
