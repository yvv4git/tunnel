package infrastructure

import (
	"context"
	"fmt"
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

func (s *ServerTCP) Start(ctx context.Context) error {
	addr, err := createAddrString(s.cfg.Host, s.cfg.Port)
	if err != nil {
		return fmt.Errorf("create server TCP address: %w", err)
	}

	listener, err := net.Listen(tcpProtocol, addr)
	if err != nil {
		return fmt.Errorf("create server TCP listener: %w", err)
	}
	s.listener = listener

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
		buffer := make([]byte, s.cfg.BufferSize)
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
		buffer := make([]byte, s.cfg.BufferSize)
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
