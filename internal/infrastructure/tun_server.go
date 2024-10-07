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

	return t.iface, nil
}

func (t *DeviceTUNServerBuilder) configureServerForLinux() error {
	cmd := exec.Command("ip", "link", "set", "dev", t.iface.Name(), "up")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("link up on TUN device: %w", err)
	}

	cfgDeviceTUN := t.cfg.Server.DeviceTUN
	cmd = exec.Command("ip", "addr", "add", cfgDeviceTUN.Host, "dev", t.iface.Name())
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("setup ip address on TUN device: %w", err)
	}

	cmd = exec.Command("ip", "route", "add", cfgDeviceTUN.Route, "dev", t.iface.Name())
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("setup ip route on TUN device: %w", err)
	}

	return nil
}

func (t *DeviceTUNServerBuilder) configureServerForMacOS() error {
	// TODO: implement
	return nil
}
