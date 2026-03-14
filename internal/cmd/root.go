package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "podcastindex",
	Short: "A CLI for the PodcastIndex API",
	Long:  "podcastindex-cli is a command-line interface for interacting with the PodcastIndex.org API.",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("api-key", "k", "", "PodcastIndex API key")
	rootCmd.PersistentFlags().StringP("api-secret", "s", "", "PodcastIndex API secret")
	rootCmd.PersistentFlags().StringP("output", "o", "table", "Output format: table, json, plain")

	viper.BindPFlag("api_key", rootCmd.PersistentFlags().Lookup("api-key"))
	viper.BindPFlag("api_secret", rootCmd.PersistentFlags().Lookup("api-secret"))
	viper.BindPFlag("output_format", rootCmd.PersistentFlags().Lookup("output"))

	viper.SetEnvPrefix("PODCASTINDEX")
	viper.AutomaticEnv()
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	configDir := configDirPath()
	viper.AddConfigPath(configDir)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Fprintf(os.Stderr, "Warning: error reading config file: %v\n", err)
		}
	}
}

func configDirPath() string {
	if xdg := os.Getenv("XDG_CONFIG_HOME"); xdg != "" {
		return filepath.Join(xdg, "podcastindex-cli")
	}
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "podcastindex-cli")
}
