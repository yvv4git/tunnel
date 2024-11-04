package application

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/yvv4git/tunnel/internal/domain/service"
	"github.com/yvv4git/tunnel/internal/infrastructure"
	"github.com/yvv4git/tunnel/internal/infrastructure/speedtest"
)

type Speedtest struct {
	application
	cfg     infrastructure.Config
	appType string
}

func NewSpeedtest(log *slog.Logger, cfg infrastructure.Config, appType string) *Speedtest {
	s := &Speedtest{
		application: application{
			log: log,
		},
		cfg:     cfg,
		appType: appType,
	}

	s.app = s
	return s
}

func (s *Speedtest) start(ctx context.Context) error {
	s.log.Info(fmt.Sprintf("Starting %s application", s.appType))
	defer s.log.Info(fmt.Sprintf("Shutting down %s application", s.appType))

	if s.appType == "server" {
		return s.startServer(ctx)
	}

	return s.startClient(ctx)
}

func (s *Speedtest) startServer(ctx context.Context) error {
	serverTCP := speedtest.NewServerTCP(s.log, s.cfg.SpeedTest.TCPServerSpeedTest)
	svc := service.NewSpeedtestServer(serverTCP)

	if err := svc.Processing(ctx); err != nil {
		return fmt.Errorf("start server: %w", err)
	}

	return nil
}

func (s *Speedtest) startClient(ctx context.Context) error {
	clientTCP := speedtest.NewClientTCP(s.log, s.cfg.SpeedTest.TCPClientSpeedTest)
	svc := service.NewSpeedtestClient(clientTCP)

	if err := svc.Processing(ctx); err != nil {
		return fmt.Errorf("start client: %w", err)
	}

	return nil
}
