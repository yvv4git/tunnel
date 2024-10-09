package infrastructure

import (
	"context"
	"fmt"
	"io"
	"net"

	"github.com/songgao/water"
)

type ClientTCP struct {
	cfg       Client
	tunDevice *water.Interface
	conn      net.Conn
}

func NewClientTCP(cfg Client, tunDevice *water.Interface) *ClientTCP {
	return &ClientTCP{
		cfg:       cfg,
		tunDevice: tunDevice,
	}
}

func (c *ClientTCP) Start(ctx context.Context) error {
	addr, err := createAddrString(c.cfg.ServerHost, c.cfg.ServerPort)
	if err != nil {
		return fmt.Errorf("create client TCP address: %w", err)
	}

	conn, err := net.Dial(tcpProtocol, addr)
	if err != nil {
		return fmt.Errorf("create client TCP connection: %w", err)
	}
	c.conn = conn

	go c.tunToTCP(ctx)
	go c.tcpToTun(ctx)

	<-ctx.Done()
	return ctx.Err()
}

func (c *ClientTCP) tunToTCP(ctx context.Context) {
	buffer := make([]byte, c.cfg.BufferSize)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			n, err := c.tunDevice.Read(buffer)
			if err != nil {
				if err == io.EOF {
					return
				}
				fmt.Printf("reading from tun device: %v\n", err)
				continue
			}

			_, err = c.conn.Write(buffer[:n])
			if err != nil {
				fmt.Printf("writing to TCP connection: %v\n", err)
				continue
			}
		}
	}
}

func (c *ClientTCP) tcpToTun(ctx context.Context) {
	buffer := make([]byte, c.cfg.BufferSize)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			n, err := c.conn.Read(buffer)
			if err != nil {
				if err == io.EOF {
					return
				}
				// TODO: Use logger
				fmt.Printf("reading from TCP connection: %v\n", err)
				continue
			}

			_, err = c.tunDevice.Write(buffer[:n])
			if err != nil {
				fmt.Printf("writing to tun device: %v\n", err)
				continue
			}
		}
	}
}

func (c *ClientTCP) Close() error {
	if c.conn != nil {
		if err := c.conn.Close(); err != nil {
			return fmt.Errorf("close TCP connection: %w", err)
		}
	}

	return nil
}
