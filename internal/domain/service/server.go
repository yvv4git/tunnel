package service

import (
	"context"
	"fmt"
	"net"

	"github.com/songgao/water"
	"github.com/yvv4git/tunnel/internal/infrastructure"
)

type Server struct {
	cfg       infrastructure.Server
	tunDevice *water.Interface
	listener  net.Listener
}

func NewServer(cfg infrastructure.Server, tunDevice *water.Interface, listener net.Listener) *Server {
	return &Server{
		cfg:       cfg,
		tunDevice: tunDevice,
		listener:  listener,
	}
}

func (s *Server) Start(ctx context.Context) error {
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

func (s *Server) handleConnection(ctx context.Context, conn net.Conn) {
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

func (s *Server) Close() error {
	if err := s.tunDevice.Close(); err != nil {
		return fmt.Errorf("close tun device: %w", err)
	}

	if err := s.listener.Close(); err != nil {
		return fmt.Errorf("close listener: %w", err)
	}

	return nil
}
