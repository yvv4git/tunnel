package infrastructure

import (
	"context"
	"fmt"
	"net"

	"github.com/songgao/water"
	"github.com/yvv4git/tunnel/internal/infrastructure/config"
)

// ServerUDP - experimental
type ServerUDP struct {
	cfg       config.Server
	tunDevice *water.Interface
	conn      *net.UDPConn
}

func NewServerUDP(cfg config.Server, tunDevice *water.Interface) *ServerUDP {
	return &ServerUDP{
		cfg:       cfg,
		tunDevice: tunDevice,
	}
}

func (s *ServerUDP) setupListener() error {
	serverCfg := s.cfg.UDPConfig
	addr, err := createAddrString(serverCfg.Host, serverCfg.Port)
	if err != nil {
		return fmt.Errorf("create server UDP address: %w", err)
	}

	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return fmt.Errorf("resolve UDP address: %w", err)
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return fmt.Errorf("create server UDP listener: %w", err)
	}

	s.conn = conn

	return nil
}

func (s *ServerUDP) Start(ctx context.Context) error {
	if err := s.setupListener(); err != nil {
		return fmt.Errorf("create server UDP listener: %w", err)
	}

	packetChan := make(chan *packet)

	go func() {
		buffer := make([]byte, s.cfg.UDPConfig.BufferSize)
		for {
			n, addr, err := s.conn.ReadFromUDP(buffer)
			if err != nil {
				if opErr, ok := err.(*net.OpError); ok && (opErr.Timeout() || opErr.Temporary()) {
					continue
				}
				close(packetChan)
				return
			}

			// TODO: Use logger
			fmt.Printf("Client connected: %s\n", addr.String())

			packetChan <- &packet{data: buffer[:n], addr: addr}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case pkt, ok := <-packetChan:
			if !ok {
				return fmt.Errorf("listener closed")
			}
			go s.handlePacket(ctx, pkt)
		}
	}
}

type packet struct {
	data []byte
	addr *net.UDPAddr
}

func (s *ServerUDP) handlePacket(ctx context.Context, pkt *packet) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in handlePacket", r)
		}
	}()

	//fromConn := make(chan []byte)
	fromTun := make(chan []byte)

	go func() {
		buffer := make([]byte, s.cfg.UDPConfig.BufferSize)
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
		case data, ok := <-fromTun:
			if !ok {
				return
			}

			_, err := s.conn.WriteToUDP(data, pkt.addr)
			if err != nil {
				fmt.Println("writing to connection:", err)
				return
			}
		default:
			_, err := s.tunDevice.Write(pkt.data)
			if err != nil {
				fmt.Println("writing to tun device:", err)
				return
			}
			return
		}
	}
}

func (s *ServerUDP) Close() error {
	if err := s.conn.Close(); err != nil {
		return fmt.Errorf("close connection: %w", err)
	}

	return nil
}
