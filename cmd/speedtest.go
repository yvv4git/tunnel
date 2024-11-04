/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log/slog"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yvv4git/tunnel/internal/application"
	"github.com/yvv4git/tunnel/internal/infrastructure"
	"github.com/yvv4git/tunnel/internal/infrastructure/config"
)

// speedtestCmd represents the speedtest command
var speedtestCmd = &cobra.Command{
	Use:   "speedtest [type]",
	Short: "Run a network speed test",
	Long: `Run a network speed test to measure the upload and download speeds of your connection.
This command uses the internal infrastructure to perform the test and logs the results.

Example usage:
  $ tunnel speedtest server
  $ tunnel speedtest client

This will initiate the speed test and display the results in the terminal.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log := infrastructure.NewDefaultLogger()
		appType := args[0]

		if appType != "server" && appType != "client" {
			log.Error("Invalid test type. Use 'upload' or 'download'.")
			return
		}

		var config config.Config
		if err := viper.Unmarshal(&config); err != nil {
			log.Error("unmarshalling config", slog.Any("error", err))
			return
		}

		app := application.NewSpeedtest(log, config, appType)
		if err := app.Start(); err != nil {
			log.Error("starting speedtest application", slog.Any("error", err))
		}
	},
}

func init() {
	rootCmd.AddCommand(speedtestCmd)
}
