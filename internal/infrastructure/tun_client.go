package infrastructure

import (
	"fmt"
	"log/slog"
	"os/exec"

	"github.com/songgao/water"
)

type DeviceTUNClientBuilder struct {
	cfg   Config
	log   *slog.Logger
	iface *water.Interface
}

func NewDeviceTUNClientBuilder(cfg Config, log *slog.Logger) (*DeviceTUNClientBuilder, error) {
	iface, err := water.New(water.Config{
		DeviceType: water.TUN,
	})
	if err != nil {
		return nil, fmt.Errorf("create client tun device: %w", err)
	}
	log.Info("device tun client created", iface.Name())

	return &DeviceTUNClientBuilder{
		cfg:   cfg,
		log:   log,
		iface: iface,
	}, nil
}

func (t *DeviceTUNClientBuilder) Build(platform Platform) (*water.Interface, error) {
	switch platform {
	case PlatformLinux:
		if err := t.configureClientForLinux(); err != nil {
			return t.iface, err
		}
		return t.iface, nil
	case PlatformMacOC:
		if err := t.configureClientForMacOS(); err != nil {
			return t.iface, err
		}
		return t.iface, nil
	}

	return t.iface, nil
}

func (t *DeviceTUNClientBuilder) configureClientForLinux() error {
	t.log.Info("configuring Linux client")

	cmd := exec.Command("ip", "link", "set", "dev", t.iface.Name(), "up")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("link up on tun device: %w", err)
	}

	cfgDeviceTUN := t.cfg.Client.DeviceTUN
	cmd = exec.Command("ip", "addr", "add", cfgDeviceTUN.Host, "dev", t.iface.Name())
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("setup ip address on tun device: %w", err)
	}

	cmd = exec.Command("ip", "route", "add", cfgDeviceTUN.Route, "dev", t.iface.Name())
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("setup ip route on tun device: %w", err)
	}

	return nil
}

func (t *DeviceTUNClientBuilder) configureClientForMacOS() error {
	t.log.Info("configuring mac client")

	cfgDeviceTUN := t.cfg.Client.DeviceTUN
	cfgServerTUN := t.cfg.Server.DeviceTUN

	// Assign IP address to the interface and bring it up
	cmd := exec.Command("sudo", "ifconfig", t.iface.Name(), cfgDeviceTUN.Host, cfgServerTUN.Host, "up")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("assign IP address and bring up tun device: %w", err)
	}

	// Add route to the interface
	cmd = exec.Command("sudo", "route", "add", "-net", cfgDeviceTUN.Route, "-interface", t.iface.Name())
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("add route to tun device: %w", err)
	}

	return nil
}
