package speedtest

import (
	"context"
	"crypto/rand"
	"fmt"
	"log/slog"
	"net"

	"github.com/yvv4git/tunnel/internal/infrastructure/config"
	"github.com/yvv4git/tunnel/internal/utils"
)

type ClientTCP struct {
	logger *slog.Logger
	cfg    config.TCPClientSpeedTest
}

func NewClientTCP(logger *slog.Logger, cfg config.TCPClientSpeedTest) *ClientTCP {
	return &ClientTCP{
		logger: logger,
		cfg:    cfg,
	}
}

func (c *ClientTCP) Start(ctx context.Context) error {
	clientAddr, err := utils.FormatAddrString(c.cfg.ServerHost, c.cfg.ServerPort)
	if err != nil {
		return fmt.Errorf("format client address: %w", err)
	}

	conn, err := net.Dial("tcp", clientAddr)
	if err != nil {
		return fmt.Errorf("failed to connect to TCP server: %w", err)
	}
	defer conn.Close()

	c.logger.Info("TCP client connected", "address", clientAddr)

	go func() {
		<-ctx.Done()
		conn.Close()
	}()

	buf := make([]byte, 1024)
	randomData := make([]byte, 1024) // TODO: use config

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			// Generate random data
			_, err = rand.Read(randomData)
			if err != nil {
				c.logger.Error("generate random data", "error", err)
				return err
			}

			// Send random data to server
			_, err = conn.Write(randomData)
			if err != nil {
				c.logger.Error("write to connection", "error", err)
				return err
			}

			// Read data from server
			_, err = conn.Read(buf)
			if err != nil {
				c.logger.Error("read from connection", "error", err)
				return err
			}
		}
	}
}
