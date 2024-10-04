package infrastructure

import (
	"fmt"
	"net"
)

const (
	TCPServer = iota
	TCPClient
)

func NewServerTCPListener(cfg Config) (net.Listener, error) {
	listener, err := net.Listen("tcp", "0.0.0.0:12345") // TODO: make this configurable
	if err != nil {
		return nil, fmt.Errorf("create server TCP listener: %w", err)
	}

	return listener, nil
}

func NewClientTCPConnection(cfg Config) (net.Conn, error) {
	conn, err := net.Dial("tcp", "server:12345") // TODO: make this configurable
	if err != nil {
		return nil, fmt.Errorf("create client TCP connection: %w", err)
	}

	return conn, nil
}
