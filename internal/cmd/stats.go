package cmd

import (
	"github.com/spf13/cobra"

	"github.com/G10xy/podcastindex-cli/internal/formatter"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "PodcastIndex statistics",
}

var statsCurrentCmd = &cobra.Command{
	Use:   "current",
	Short: "Current index statistics",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := newClient().StatsCurrent(cmd.Context())
		if err != nil {
			return err
		}
		return output(result, []formatter.TableRow{result.Stats})
	},
}

func init() {
	statsCmd.AddCommand(statsCurrentCmd)
	rootCmd.AddCommand(statsCmd)
}
