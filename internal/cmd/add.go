package cmd

import (
	"github.com/spf13/cobra"

	"github.com/G10xy/podcastindex-cli/internal/client"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add podcasts to the PodcastIndex (requires write-enabled API key)",
}

var addByFeedURLCmd = &cobra.Command{
	Use:   "byfeedurl",
	Short: "Add podcast by feed URL",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := newClient().AddByFeedURL(cmd.Context(), client.AddByFeedURLOptions{
			URL:      mustString(cmd, "url"),
			CHash:    mustString(cmd, "chash"),
			ItunesID: mustInt(cmd, "itunesid"),
		})
		if err != nil {
			return err
		}
		return output(result, nil)
	},
}

var addByItunesIDCmd = &cobra.Command{
	Use:   "byitunesid",
	Short: "Add podcast by iTunes ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := newClient().AddByItunesID(cmd.Context(), mustInt(cmd, "id"))
		if err != nil {
			return err
		}
		return output(result, nil)
	},
}

func init() {
	addByFeedURLCmd.Flags().String("url", "", "Podcast feed URL (required)")
	addByFeedURLCmd.MarkFlagRequired("url")
	addByFeedURLCmd.Flags().String("chash", "", "MD5 hash for duplicate checking")
	addByFeedURLCmd.Flags().Int("itunesid", 0, "Associate this iTunes ID if none exists")

	addByItunesIDCmd.Flags().Int("id", 0, "iTunes ID (required)")
	addByItunesIDCmd.MarkFlagRequired("id")

	addCmd.AddCommand(addByFeedURLCmd, addByItunesIDCmd)
	rootCmd.AddCommand(addCmd)
}
