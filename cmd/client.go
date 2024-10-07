/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log/slog"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yvv4git/tunnel/internal/application"
	"github.com/yvv4git/tunnel/internal/infrastructure"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Start the application client",
	Long: `Start the application client with the specified configuration.

This command initializes and runs the client component of the application. 
It reads the configuration from the specified file or default location and connects to the server on the specified host and port.

Example usage:
  app client
  app client --config /path/to/config.toml

The client command will load the configuration and start the client, allowing it to interact with the server.`,
	Run: func(cmd *cobra.Command, args []string) {
		log := infrastructure.NewDefaultLogger()

		var config infrastructure.Config
		if err := viper.Unmarshal(&config); err != nil {
			log.Error("unmarshalling config", slog.Any("error", err))
			return
		}

		spew.Dump(config)
		app := application.NewClient(log, config)
		if err := app.Start(); err != nil {
			log.Error("start client application", slog.Any("error", err))
		}
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
