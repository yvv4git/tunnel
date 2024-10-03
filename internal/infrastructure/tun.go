package infrastructure

import (
	"fmt"
	"os/exec"

	"github.com/songgao/water"
)

type Platform string

const (
	PlatformLinux Platform = "linux"
	PlatformMacOC Platform = "macos"
)

type NodeType string

const (
	NodeTypeServer NodeType = "server"
	NodeTypeClient NodeType = "client"
)

type TUNDeviceBuilder struct {
	iface *water.Interface
}

func NewTUNDeviceBuilder() (*TUNDeviceBuilder, error) {
	iface, err := water.New(water.Config{
		DeviceType: water.TUN,
	})
	if err != nil {
		return nil, fmt.Errorf("create server TUN device: %w", err)
	}

	return &TUNDeviceBuilder{
		iface: iface,
	}, nil
}

func (t *TUNDeviceBuilder) Build(platform Platform, nodeType NodeType) *water.Interface {
	switch {
	case (platform == PlatformLinux) && (nodeType == NodeTypeServer):
		t.configureServerForLinux()
	case (platform == PlatformLinux) && (nodeType == NodeTypeClient):
		t.configureClientForLinux()
	case (platform == PlatformMacOC) && (nodeType == NodeTypeServer):
		t.configureServerForMacOS()
	case (platform == PlatformMacOC) && (nodeType == NodeTypeClient):
		t.configureClientForMacOS()
	}

	return t.iface
}

func (t *TUNDeviceBuilder) configureServerForLinux() error {
	cmd := exec.Command("ip", "link", "set", "dev", t.iface.Name(), "up")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("link up on TUN device: %w", err)
	}

	cmd = exec.Command("ip", "addr", "add", "10.0.0.1/24", "dev", t.iface.Name()) // TODO: make this configurable
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("setup ip address on TUN device: %w", err)
	}

	return nil
}

func (t *TUNDeviceBuilder) configureClientForLinux() error {
	cmd := exec.Command("ip", "link", "set", "dev", t.iface.Name(), "up")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("link up on TUN device: %w", err)
	}

	cmd = exec.Command("ip", "addr", "add", "10.0.0.2/24", "dev", t.iface.Name()) // TODO: make this configurable
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("setup ip address on TUN device: %w", err)
	}

	cmd = exec.Command("ip", "route", "add", "10.0.0.0/24", "dev", t.iface.Name()) // TODO: make this configurable
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("setup ip route on TUN device: %w", err)
	}

	return nil
}

func (t *TUNDeviceBuilder) configureServerForMacOS() error {
	// TODO: implement
	return nil
}

func (t *TUNDeviceBuilder) configureClientForMacOS() error {
	// TODO: implement
	return nil
}
