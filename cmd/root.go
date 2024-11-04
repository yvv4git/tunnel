/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yvv4git/tunnel/internal/infrastructure"
	"github.com/yvv4git/tunnel/internal/infrastructure/config"
)

// rootCmd represents the root command
var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Loads the configuration",
	Long: `This command initializes the application by loading the necessary configuration files.
It serves as the entry point for the application and sets up the environment for further operations.

Example usage:
  app
  app --config /path/to/config.yaml

This command does not perform any additional actions beyond loading the configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		log := infrastructure.NewDefaultLogger()

		var config config.Config
		if err := viper.Unmarshal(&config); err != nil {
			log.Error("unmarshalling config", slog.Any("error", err))
			return
		}

		spew.Dump(config)
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringP("config", "c", "./configs/config.yaml", "Path to the config file")
	if err := viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config")); err != nil {
		log.Fatalf("Error binding config flag: %s", err.Error())
	}
}

func initConfig() {
	configFile := viper.GetString("config")
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err)
		os.Exit(1)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
