package infrastructure

import (
	"fmt"
	"net"
)

const (
	tcpProtocol = "tcp"
)

func NewServerTCPListener(cfg Config) (net.Listener, error) {
	addr, err := createAddrString(cfg.Server.Host, cfg.Server.Port)
	if err != nil {
		return nil, fmt.Errorf("create server TCP address: %w", err)
	}

	listener, err := net.Listen(tcpProtocol, addr)
	if err != nil {
		return nil, fmt.Errorf("create server TCP listener: %w", err)
	}

	return listener, nil
}

func NewClientTCPConnection(cfg Config) (net.Conn, error) {
	addr, err := createAddrString(cfg.Client.ServerHost, cfg.Client.ServerPort)
	if err != nil {
		return nil, fmt.Errorf("create client TCP address: %w", err)
	}

	conn, err := net.Dial(tcpProtocol, addr)
	if err != nil {
		return nil, fmt.Errorf("create client TCP connection: %w", err)
	}

	return conn, nil
}

func createAddrString(host string, port uint16) (string, error) {
	if host == "" {
		return "", ErrInvalidHost
	}

	if port == 0 {
		return "", ErrInvalidPort
	}

	return fmt.Sprintf("%s:%d", host, port), nil
}
