package infrastructure

import (
	"fmt"
	"os/exec"

	"github.com/songgao/water"
)

type DeviceTUNClientBuilder struct {
	cfg   Config
	iface *water.Interface
}

func NewDeviceTUNClientBuilder(cfg Config) (*DeviceTUNClientBuilder, error) {
	iface, err := water.New(water.Config{
		DeviceType: water.TUN,
	})
	if err != nil {
		return nil, fmt.Errorf("create client TUN device: %w", err)
	}

	return &DeviceTUNClientBuilder{
		cfg:   cfg,
		iface: iface,
	}, nil
}

func (t *DeviceTUNClientBuilder) Build(platform Platform) (*water.Interface, error) {
	switch platform {
	case PlatformLinux:
		t.configureClientForLinux()
		return t.iface, nil
	case PlatformMacOC:
		t.configureClientForMacOS()
		return t.iface, nil
	}

	return t.iface, nil
}

func (t *DeviceTUNClientBuilder) configureClientForLinux() error {
	cmd := exec.Command("ip", "link", "set", "dev", t.iface.Name(), "up")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("link up on TUN device: %w", err)
	}

	cfgDeviceTUN := t.cfg.Client.DeviceTUN
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

func (t *DeviceTUNClientBuilder) configureClientForMacOS() error {
	// TODO: implement
	return nil
}
