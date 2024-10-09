package service

import (
	"context"
	"fmt"
	"io"
	"net"

	"github.com/songgao/water"
	"github.com/yvv4git/tunnel/internal/infrastructure"
)

type Client struct {
	cfg       infrastructure.Client
	tunDevice *water.Interface
	listener  net.Listener
	conn      net.Conn
}

func NewClient(cfg infrastructure.Client, tunDevice *water.Interface, listener net.Listener) *Client {
	return &Client{
		cfg:       cfg,
		tunDevice: tunDevice,
		listener:  listener,
	}
}

func (c *Client) Start(ctx context.Context) error {
	go c.tunToTCP(ctx)

	go c.tcpToTun(ctx)

	return nil
}

func (c *Client) tunToTCP(ctx context.Context) {
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

func (c *Client) tcpToTun(ctx context.Context) {
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

func (c *Client) Close() error {
	if err := c.tunDevice.Close(); err != nil {
		return fmt.Errorf("close tun device: %w", err)
	}

	if err := c.listener.Close(); err != nil {
		return fmt.Errorf("close listener: %w", err)
	}

	if c.conn != nil {
		if err := c.conn.Close(); err != nil {
			return fmt.Errorf("close TCP connection: %w", err)
		}
	}

	return nil
}
