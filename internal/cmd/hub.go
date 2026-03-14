package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/G10xy/podcastindex-cli/internal/client"
)

var hubCmd = &cobra.Command{
	Use:   "hub",
	Short: "Hub notification commands",
}

var hubPubNotifyCmd = &cobra.Command{
	Use:   "pubnotify",
	Short: "Notify index of feed change",
	RunE: func(cmd *cobra.Command, args []string) error {
		id := mustInt(cmd, "id")
		u := mustString(cmd, "url")
		if id == 0 && u == "" {
			return fmt.Errorf("--id or --url is required")
		}
		result, err := newClient().HubPubNotify(cmd.Context(), client.HubPubNotifyOptions{ID: id, URL: u})
		if err != nil {
			return err
		}
		return output(result, nil)
	},
}

func init() {
	hubPubNotifyCmd.Flags().Int("id", 0, "PodcastIndex Feed ID")
	hubPubNotifyCmd.Flags().String("url", "", "Podcast feed URL")
	hubCmd.AddCommand(hubPubNotifyCmd)
	rootCmd.AddCommand(hubCmd)
}
