package infrastructure

import (
	"fmt"
	"net"
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
