package application

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/yvv4git/tunnel/internal/domain/service"
	"github.com/yvv4git/tunnel/internal/infrastructure/config"
	"github.com/yvv4git/tunnel/internal/infrastructure/direct"
)

type Server struct {
	application
	cfg config.Config
}

func NewServer(log *slog.Logger, cfg config.Config) *Server {
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

	tunDeviceBuilder, err := direct.NewDeviceTUNServerBuilder(s.cfg, s.log)
	if err != nil {
		return fmt.Errorf("create TUN device builder: %w", err)
	}

	tunDeviceCfg := s.cfg.DirectConnection.Server.DeviceTUN
	currentPlatform := direct.Platform(tunDeviceCfg.Platform)
	tunDevice, err := tunDeviceBuilder.Build(currentPlatform)
	if err != nil {
		return fmt.Errorf("build TUN device: %w", err)
	}
	defer tunDevice.Close()

	channelServerBuilder := direct.NewChannelServerBuilder(s.log, s.cfg, tunDevice)
	channelServer, err := channelServerBuilder.Build(s.cfg.DirectConnection.Server.ChannelType)
	if err != nil {
		return fmt.Errorf("build server TCP: %w", err)
	}
	defer channelServer.Close()

	metricsWebServerCfg := s.cfg.DirectConnection.Server.TCPConfig.Metrics
	direct.StartMetricsWebServer(metricsWebServerCfg)

	svc := service.NewServer(channelServer)
	if err = svc.Processing(ctx); err != nil {
		return fmt.Errorf("start server: %w", err)
	}

	return nil
}
