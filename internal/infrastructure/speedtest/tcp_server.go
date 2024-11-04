package speedtest

import (
	"context"
	"crypto/rand"
	"fmt"
	"log/slog"
	"net"

	"github.com/yvv4git/tunnel/internal/infrastructure"
	"github.com/yvv4git/tunnel/internal/utils"
)

type ServerTCP struct {
	logger   *slog.Logger
	cfg      infrastructure.TCPServerSpeedTest
	listener net.Listener
}

func NewServerTCP(logger *slog.Logger, cfg infrastructure.TCPServerSpeedTest) *ServerTCP {
	return &ServerTCP{
		logger: logger,
		cfg:    cfg,
	}
}

func (s *ServerTCP) Start(ctx context.Context) error {
	var err error
	serverAddr, err := utils.FormatAddrString(s.cfg.Host, s.cfg.Port)
	if err != nil {
		return fmt.Errorf("format server address: %w", err)
	}

	s.listener, err = net.Listen("tcp", serverAddr)
	if err != nil {
		return fmt.Errorf("failed to start TCP server: %w", err)
	}

	s.logger.Info("TCP server started", "address", serverAddr)

	go func() {
		<-ctx.Done()
		s.Close()
	}()

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			select {
			case <-ctx.Done():
				return nil
			default:
				s.logger.Error("failed to accept connection", "error", err)
				continue
			}
		}

		go s.handleConnection(ctx, conn)
	}
}

func (s *ServerTCP) handleConnection(ctx context.Context, conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	randomData := make([]byte, 1024)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			// Read data from client
			_, err := conn.Read(buf)
			if err != nil {
				s.logger.Error("read from connection", "error", err)
				return
			}

			// Generate random data
			_, err = rand.Read(randomData)
			if err != nil {
				s.logger.Error("generate random data", "error", err)
				return
			}

			// Send random data to client
			_, err = conn.Write(randomData)
			if err != nil {
				s.logger.Error("write to connection", "error", err)
				return
			}
		}
	}
}

func (s *ServerTCP) Close() error {
	if s.listener != nil {
		return s.listener.Close()
	}

	return nil
}
