package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage CLI configuration",
}

var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Display current configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		settings := map[string]interface{}{
			"api_key":       viper.GetString("api_key"),
			"api_secret":    maskSecret(viper.GetString("api_secret")),
			"output_format": viper.GetString("output_format"),
		}
		for k, v := range settings {
			fmt.Fprintf(os.Stdout, "%s: %v\n", k, v)
		}
		return nil
	},
}

var configSetCmd = &cobra.Command{
	Use:   "set <key> <value>",
	Short: "Set a configuration value",
	Long:  "Supported keys: api_key, api_secret, output_format",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		key, value := args[0], args[1]

		switch key {
		case "api_key", "api_secret", "output_format":
		default:
			return fmt.Errorf("unsupported config key %q. Supported: api_key, api_secret, output_format", key)
		}

		viper.Set(key, value)

		configDir := configDirPath()
		if err := os.MkdirAll(configDir, 0700); err != nil {
			return fmt.Errorf("creating config directory: %w", err)
		}

		configFile := filepath.Join(configDir, "config.yaml")
		if err := viper.WriteConfigAs(configFile); err != nil {
			return fmt.Errorf("writing config file: %w", err)
		}

		fmt.Fprintf(os.Stdout, "Set %s in %s\n", key, configFile)
		return nil
	},
}

func init() {
	configCmd.AddCommand(configShowCmd)
	configCmd.AddCommand(configSetCmd)
	rootCmd.AddCommand(configCmd)
}

func maskSecret(s string) string {
	if len(s) <= 4 {
		return "****"
	}
	return s[:2] + "****" + s[len(s)-2:]
}
