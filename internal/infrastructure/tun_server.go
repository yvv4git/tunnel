package infrastructure

import (
	"fmt"
	"os/exec"

	"github.com/songgao/water"
)

type DeviceTUNServerBuilder struct {
	cfg   Config
	iface *water.Interface
}

func NewDeviceTUNServerBuilder(cfg Config) (*DeviceTUNServerBuilder, error) {
	iface, err := water.New(water.Config{
		DeviceType: water.TUN,
	})
	if err != nil {
		return nil, fmt.Errorf("create server TUN device: %w", err)
	}

	return &DeviceTUNServerBuilder{
		cfg:   cfg,
		iface: iface,
	}, nil
}

func (t *DeviceTUNServerBuilder) Build(platform Platform) (*water.Interface, error) {
	switch platform {
	case PlatformLinux:
		t.configureServerForLinux()
		return t.iface, nil
	case PlatformMacOC:
		t.configureServerForMacOS()
		return t.iface, nil
	}

	return t.iface, ErrInvalidPlatform
}

func (t *DeviceTUNServerBuilder) configureServerForLinux() error {
	cfgServerTUN := t.cfg.Server.DeviceTUN
	cfgClientTUN := t.cfg.Client.DeviceTUN

	// Bring the interface up
	cmd := exec.Command("sudo", "ip", "link", "set", "dev", t.iface.Name(), "up")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("bring up TUN device: %w", err)
	}

	// Assign IP address to the interface
	cmd = exec.Command("sudo", "ip", "addr", "add", cfgServerTUN.Host+"/32", "peer", cfgClientTUN.Host, "dev", t.iface.Name())
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("assign IP address to TUN device: %w", err)
	}

	// Add route to the interface
	cmd = exec.Command("sudo", "ip", "route", "add", cfgServerTUN.Route, "dev", t.iface.Name())
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("add route to TUN device: %w", err)
	}

	return nil
}

func (t *DeviceTUNServerBuilder) configureServerForMacOS() error {
	cfgDeviceTUN := t.cfg.Server.DeviceTUN
	cfgClientTUN := t.cfg.Client.DeviceTUN

	// Bring the interface up
	cmd := exec.Command("sudo", "ifconfig", t.iface.Name(), "up")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("bring up TUN device: %w", err)
	}

	// Assign IP address to the interface
	cmd = exec.Command("sudo", "ifconfig", t.iface.Name(), cfgDeviceTUN.Host, cfgClientTUN.Host, "netmask", "255.255.255.255")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("assign IP address to TUN device: %w", err)
	}

	// Add route to the interface
	cmd = exec.Command("sudo", "route", "add", "-net", cfgDeviceTUN.Route, "-interface", t.iface.Name())
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("add route to TUN device: %w", err)
	}

	return nil
}
