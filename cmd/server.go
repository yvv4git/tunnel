/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log/slog"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yvv4git/tunnel/internal/infrastructure"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the application server",
	Long: `Start the application server with the specified configuration.

This command initializes and runs the server component of the application. It reads the configuration from the specified file or default location and starts the server on the specified host and port.

Example usage:
  app server
  app server --config /path/to/config.toml

The server command will load the configuration and start the server, making the application available for incoming requests.`,
	Run: func(cmd *cobra.Command, args []string) {
		log := infrastructure.NewDefaultLogger()

		var config infrastructure.Config
		err := viper.Unmarshal(&config)
		if err != nil {
			log.Error("unmarshalling config", slog.Any("error", err))
			return
		}

		log.Info("config loaded", slog.Any("config", config))
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
