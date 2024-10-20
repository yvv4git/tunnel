package infrastructure

import (
	"context"
	"fmt"
	"net"

	"github.com/songgao/water"
)

// ClientUDP - experimental, not work now
type ClientUDP struct {
	cfg       Client
	tunDevice *water.Interface
	conn      *net.UDPConn
}

func NewClientUDP(cfg Client, tunDevice *water.Interface) *ClientUDP {
	return &ClientUDP{
		cfg:       cfg,
		tunDevice: tunDevice,
	}
}

func (c *ClientUDP) Start(ctx context.Context) error {
	if err := c.setupConn(); err != nil {
		return fmt.Errorf("create client UDP connection: %w", err)
	}

	go c.tunToUDP(ctx)
	go c.udpToTun(ctx)

	<-ctx.Done()
	return ctx.Err()
}

func (c *ClientUDP) setupConn() error {
	clientCfg := c.cfg.UDPConfig
	addr, err := createAddrString(clientCfg.ServerHost, clientCfg.ServerPort)
	if err != nil {
		return fmt.Errorf("create client UDP address: %w", err)
	}

	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return fmt.Errorf("resolve UDP address: %w", err)
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		return fmt.Errorf("create client UDP connection: %w", err)
	}
	c.conn = conn

	return nil
}

func (c *ClientUDP) tunToUDP(ctx context.Context) {
	buffer := make([]byte, c.cfg.UDPConfig.BufferSize)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			n, err := c.tunDevice.Read(buffer)
			if err != nil {
				continue
			}

			_, err = c.conn.Write(buffer[:n])
			if err != nil {
				continue
			}
		}
	}
}

func (c *ClientUDP) udpToTun(ctx context.Context) {
	buffer := make([]byte, c.cfg.UDPConfig.BufferSize)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			n, _, err := c.conn.ReadFromUDP(buffer)
			if err != nil {
				continue
			}

			_, err = c.tunDevice.Write(buffer[:n])
			if err != nil {
				continue
			}
		}
	}
}

func (c *ClientUDP) Close() error {
	if c.conn != nil {
		if err := c.conn.Close(); err != nil {
			return fmt.Errorf("close UDP connection: %w", err)
		}
	}

	return nil
}
