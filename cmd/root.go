/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		fmt.Println("Configuration loaded")
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringP("config", "c", "config.toml", "Path to the config file")
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
		viper.SetConfigType("toml")
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
