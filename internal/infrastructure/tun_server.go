package infrastructure

import (
	"fmt"
	"log/slog"
	"os/exec"

	"github.com/songgao/water"
)

type DeviceTUNServerBuilder struct {
	cfg   Config
	log   *slog.Logger
	iface *water.Interface
}

func NewDeviceTUNServerBuilder(cfg Config, log *slog.Logger) (*DeviceTUNServerBuilder, error) {
	iface, err := water.New(water.Config{
		DeviceType: water.TUN,
	})
	if err != nil {
		return nil, fmt.Errorf("create server tun device: %w", err)
	}
	log.Info("create tun device", slog.String("device", iface.Name()))

	return &DeviceTUNServerBuilder{
		cfg:   cfg,
		log:   log,
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
	t.log.Info("configure server for linux")

	cfgServerTUN := t.cfg.Server.DeviceTUN
	cfgClientTUN := t.cfg.Client.DeviceTUN

	// Bring the interface up
	cmd := exec.Command("ip", "link", "set", "dev", t.iface.Name(), "up")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("bring up tun device: %w", err)
	}

	// Assign IP address to the interface
	cmd = exec.Command("ip", "addr", "add", cfgServerTUN.Host+"/32", "peer", cfgClientTUN.Host, "dev", t.iface.Name())
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("assign IP address to tun device: %w", err)
	}

	// Add route to the interface
	cmd = exec.Command("ip", "route", "add", cfgServerTUN.Route, "dev", t.iface.Name())
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("add route to tun device: %w", err)
	}

	return nil
}

func (t *DeviceTUNServerBuilder) configureServerForMacOS() error {
	t.log.Info("configure server for macos")

	cfgServerTUN := t.cfg.Server.DeviceTUN
	cfgClientTUN := t.cfg.Client.DeviceTUN

	// Bring the interface up
	cmd := exec.Command("ifconfig", t.iface.Name(), "up")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("bring up tun device: %w", err)
	}

	// Assign IP address to the interface
	cmd = exec.Command("ifconfig", t.iface.Name(), cfgServerTUN.Host, cfgClientTUN.Host, "netmask", "255.255.255.255")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("assign IP address to tun device: %w", err)
	}

	// Add route to the interface
	cmd = exec.Command("route", "add", "-net", cfgServerTUN.Route, "-interface", t.iface.Name())
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("add route to tun device: %w", err)
	}

	return nil
}
