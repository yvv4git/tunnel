package infrastructure

import (
	"fmt"
	"net"
)

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
