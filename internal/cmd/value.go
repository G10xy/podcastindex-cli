package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var valueCmd = &cobra.Command{
	Use:   "value",
	Short: "Query value information from the PodcastIndex",
}

var valueByFeedIDCmd = &cobra.Command{
	Use:   "byfeedid",
	Short: "Value info by Feed ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		id := mustInt(cmd, "id")
		if id == 0 {
			return fmt.Errorf("--id is required")
		}
		result, err := newClient().ValueByFeedID(cmd.Context(), id)
		if err != nil {
			return err
		}
		return output(result, toRows(result.Value.Destinations))
	},
}

var valueByFeedURLCmd = &cobra.Command{
	Use:   "byfeedurl",
	Short: "Value info by feed URL",
	RunE: func(cmd *cobra.Command, args []string) error {
		u := mustString(cmd, "url")
		if u == "" {
			return fmt.Errorf("--url is required")
		}
		result, err := newClient().ValueByFeedURL(cmd.Context(), u)
		if err != nil {
			return err
		}
		return output(result, toRows(result.Value.Destinations))
	},
}

var valueByPodcastGUIDCmd = &cobra.Command{
	Use:   "bypodcastguid",
	Short: "Value info by Podcast GUID",
	RunE: func(cmd *cobra.Command, args []string) error {
		guid := mustString(cmd, "guid")
		if guid == "" {
			return fmt.Errorf("--guid is required")
		}
		result, err := newClient().ValueByPodcastGUID(cmd.Context(), guid)
		if err != nil {
			return err
		}
		return output(result, toRows(result.Value.Destinations))
	},
}

var valueByEpisodeGUIDCmd = &cobra.Command{
	Use:   "byepisodeguid",
	Short: "Value info by episode GUID",
	RunE: func(cmd *cobra.Command, args []string) error {
		pg := mustString(cmd, "podcastguid")
		eg := mustString(cmd, "episodeguid")
		if pg == "" || eg == "" {
			return fmt.Errorf("--podcastguid and --episodeguid are required")
		}
		result, err := newClient().ValueByEpisodeGUID(cmd.Context(), pg, eg)
		if err != nil {
			return err
		}
		return output(result, toRows(result.Value.Destinations))
	},
}

var valueBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch episode value lookup",
	RunE: func(cmd *cobra.Command, args []string) error {
		data := mustString(cmd, "data")
		file := mustString(cmd, "file")
		if data == "" && file == "" {
			return fmt.Errorf("--data or --file is required")
		}

		var raw []byte
		if file != "" {
			var err error
			raw, err = os.ReadFile(file)
			if err != nil {
				return fmt.Errorf("reading file: %w", err)
			}
		} else {
			raw = []byte(data)
		}

		var batch map[string][]string
		if err := json.Unmarshal(raw, &batch); err != nil {
			return fmt.Errorf("parsing JSON input: %w", err)
		}

		result, err := newClient().ValueBatchByEpisodeGUID(cmd.Context(), batch)
		if err != nil {
			return err
		}
		// Batch response is a map, only JSON output makes sense
		return output(result, nil)
	},
}

func init() {
	rootCmd.AddCommand(valueCmd)

	valueByFeedIDCmd.Flags().Int("id", 0, "Feed ID (required)")
	valueByFeedURLCmd.Flags().String("url", "", "Feed URL (required)")
	valueByPodcastGUIDCmd.Flags().String("guid", "", "Podcast GUID (required)")

	valueByEpisodeGUIDCmd.Flags().String("podcastguid", "", "Podcast GUID (required)")
	valueByEpisodeGUIDCmd.Flags().String("episodeguid", "", "Episode GUID (required)")

	valueBatchCmd.Flags().String("data", "", "Inline JSON: {\"podcastGUID\":[\"epGUID1\",\"epGUID2\"]}")
	valueBatchCmd.Flags().String("file", "", "Path to JSON file with batch input")

	for _, c := range []*cobra.Command{
		valueByFeedIDCmd, valueByFeedURLCmd, valueByPodcastGUIDCmd,
		valueByEpisodeGUIDCmd, valueBatchCmd,
	} {
		valueCmd.AddCommand(c)
	}
}
