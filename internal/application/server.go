package application

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/yvv4git/tunnel/internal/domain/service"
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
	s.log.Info("Starting server application")
	defer s.log.Info("Shutting down server application")

	tunDeviceBuilder, err := infrastructure.NewDeviceTUNServerBuilder(s.cfg, s.log)
	if err != nil {
		return fmt.Errorf("create TUN device builder: %w", err)
	}

	tunDeviceCfg := s.cfg.Server.DeviceTUN
	currentPlatform := infrastructure.Platform(tunDeviceCfg.Platform)
	tunDevice, err := tunDeviceBuilder.Build(currentPlatform)
	if err != nil {
		return fmt.Errorf("build TUN device: %w", err)
	}
	defer tunDevice.Close()

	channelServerBuilder := infrastructure.NewChannelServerBuilder(s.log, s.cfg, tunDevice)
	channelServer, err := channelServerBuilder.Build(s.cfg.Server.ChannelType)
	if err != nil {
		return fmt.Errorf("build server TCP: %w", err)
	}
	defer channelServer.Close()

	svc := service.NewServer(channelServer)
	if err = svc.Processing(ctx); err != nil {
		return fmt.Errorf("start server: %w", err)
	}

	return nil
}
