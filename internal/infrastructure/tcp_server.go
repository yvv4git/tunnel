package infrastructure

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"github.com/songgao/water"
)

type ServerTCP struct {
	cfg       Server
	tunDevice *water.Interface
	listener  net.Listener
}

func NewServerTCP(cfg Server, tunDevice *water.Interface) *ServerTCP {
	return &ServerTCP{
		cfg:       cfg,
		tunDevice: tunDevice,
	}
}

func (s *ServerTCP) setupListener() error {
	serverCfg := s.cfg.TCPConfig
	addr, err := createAddrString(serverCfg.Host, serverCfg.Port)
	if err != nil {
		return fmt.Errorf("create server TCP address: %w", err)
	}

	if serverCfg.Encryption.Enabled {
		encCfg := serverCfg.Encryption
		cer, err := tls.LoadX509KeyPair(encCfg.ServerCert, encCfg.ServerKey)
		if err != nil {
			log.Fatalf("Failed to load server certificates: %v", err)
		}

		clientCAs, err := loadClientCA(encCfg.CACert)
		if err != nil {
			return fmt.Errorf("load client CA: %w", err)
		}

		config := &tls.Config{
			Certificates: []tls.Certificate{cer},
			ClientAuth:   tls.RequireAndVerifyClientCert,
			ClientCAs:    clientCAs,
		}

		listener, err := tls.Listen(tcpProtocol, addr, config)
		if err != nil {
			return fmt.Errorf("create server TLS listener: %w", err)
		}

		s.listener = listener

		return nil
	}

	listener, err := net.Listen(tcpProtocol, addr)
	if err != nil {
		return fmt.Errorf("create server TCP listener: %w", err)
	}

	s.listener = listener

	return nil
}

func loadClientCA(caFile string) (*x509.CertPool, error) {
	caCert, err := ioutil.ReadFile(caFile)
	if err != nil {
		return nil, fmt.Errorf("read CA certificate: %w", err)
	}

	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM(caCert) {
		return nil, fmt.Errorf("append CA certificate: %v", err)
	}

	return caCertPool, nil
}

func (s *ServerTCP) Start(ctx context.Context) error {
	if err := s.setupListener(); err != nil {
		return fmt.Errorf("create server TCP listener: %w", err)
	}

	connChan := make(chan net.Conn)

	go func() {
		for {
			conn, err := s.listener.Accept()
			if err != nil {
				if opErr, ok := err.(*net.OpError); ok && (opErr.Timeout() || opErr.Temporary()) {
					continue
				}
				close(connChan)
				return
			}

			// TODO: Use logger
			fmt.Printf("Client connected: %s\n", conn.RemoteAddr().String())

			connChan <- conn
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case conn, ok := <-connChan:
			if !ok {
				return fmt.Errorf("listener closed")
			}
			go s.handleConnection(ctx, conn)
		}
	}
}

func (s *ServerTCP) handleConnection(ctx context.Context, conn net.Conn) {
	defer conn.Close()

	fromConn := make(chan []byte)
	fromTun := make(chan []byte)

	go func() {
		buffer := make([]byte, s.cfg.TCPConfig.BufferSize)
		for {
			n, err := conn.Read(buffer)
			if err != nil {
				fmt.Println("reading from connection:", err)
				close(fromConn)
				return
			}
			fromConn <- buffer[:n]
		}
	}()

	go func() {
		buffer := make([]byte, s.cfg.TCPConfig.BufferSize)
		for {
			n, err := s.tunDevice.Read(buffer)
			if err != nil {
				fmt.Println("reading from tun device:", err)
				close(fromTun)
				return
			}
			fromTun <- buffer[:n]
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case data, ok := <-fromConn:
			if !ok {
				return
			}

			_, err := s.tunDevice.Write(data)
			if err != nil {
				fmt.Println("writing to tun device:", err)
				return
			}
		case data, ok := <-fromTun:
			if !ok {
				return
			}

			_, err := conn.Write(data)
			if err != nil {
				fmt.Println("writing to connection:", err)
				return
			}
		}
	}
}

func (s *ServerTCP) Close() error {
	if err := s.listener.Close(); err != nil {
		return fmt.Errorf("close listener: %w", err)
	}

	return nil
}
