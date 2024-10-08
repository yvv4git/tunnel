package infrastructure

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
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
	if err := c.setupConn(); err != nil {
		return fmt.Errorf("create client TCP connection: %w", err)
	}

	go c.tunToTCP(ctx)
	go c.tcpToTun(ctx)

	<-ctx.Done()
	return ctx.Err()
}

func (c *ClientTCP) setupConn() error {
	clientCfg := c.cfg.TCPConfig
	addr, err := createAddrString(clientCfg.ServerHost, clientCfg.ServerPort)
	if err != nil {
		return fmt.Errorf("create client TCP address: %w", err)
	}

	if clientCfg.Encryption.Enabled {
		encCfg := clientCfg.Encryption
		cer, err := tls.LoadX509KeyPair(encCfg.ClientCert, encCfg.ClientKey)
		if err != nil {
			return fmt.Errorf("load client certificates: %w", err)
		}

		caCert, err := ioutil.ReadFile(encCfg.CACert)
		if err != nil {
			return fmt.Errorf("read CA certificate: %w", err)
		}

		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		config := &tls.Config{
			Certificates: []tls.Certificate{cer},
			RootCAs:      caCertPool,
		}

		conn, err := tls.Dial(tcpProtocol, addr, config)
		if err != nil {
			return fmt.Errorf("create client TLS connection: %w", err)
		}
		c.conn = conn // setup tls conn

		return nil
	}

	conn, err := net.Dial(tcpProtocol, addr)
	if err != nil {
		return fmt.Errorf("create client TCP connection: %w", err)
	}
	c.conn = conn // setup plain conn

	return nil
}

func (c *ClientTCP) tunToTCP(ctx context.Context) {
	buffer := make([]byte, c.cfg.TCPConfig.BufferSize)

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
	buffer := make([]byte, c.cfg.TCPConfig.BufferSize)

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
