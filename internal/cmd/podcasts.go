package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/G10xy/podcastindex-cli/internal/client"
	"github.com/G10xy/podcastindex-cli/internal/formatter"
)

var podcastsCmd = &cobra.Command{
	Use:   "podcasts",
	Short: "Query podcasts from the PodcastIndex",
}

var podcastsByFeedIDCmd = &cobra.Command{
	Use:   "byfeedid",
	Short: "Get podcast by PodcastIndex Feed ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		id := mustInt(cmd, "id")
		if id == 0 {
			return fmt.Errorf("--id is required")
		}
		result, err := newClient().PodcastByFeedID(cmd.Context(), id)
		if err != nil {
			return err
		}
		return output(result, []formatter.TableRow{result.Feed})
	},
}

var podcastsByFeedURLCmd = &cobra.Command{
	Use:   "byfeedurl",
	Short: "Get podcast by feed URL",
	RunE: func(cmd *cobra.Command, args []string) error {
		u := mustString(cmd, "url")
		if u == "" {
			return fmt.Errorf("--url is required")
		}
		result, err := newClient().PodcastByFeedURL(cmd.Context(), u)
		if err != nil {
			return err
		}
		return output(result, []formatter.TableRow{result.Feed})
	},
}

var podcastsByItunesIDCmd = &cobra.Command{
	Use:   "byitunesid",
	Short: "Get podcast by iTunes ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		id := mustInt(cmd, "id")
		if id == 0 {
			return fmt.Errorf("--id is required")
		}
		result, err := newClient().PodcastByItunesID(cmd.Context(), id)
		if err != nil {
			return err
		}
		return output(result, []formatter.TableRow{result.Feed})
	},
}

var podcastsByGUIDCmd = &cobra.Command{
	Use:   "byguid",
	Short: "Get podcast by GUID",
	RunE: func(cmd *cobra.Command, args []string) error {
		guid := mustString(cmd, "guid")
		if guid == "" {
			return fmt.Errorf("--guid is required")
		}
		result, err := newClient().PodcastByGUID(cmd.Context(), guid)
		if err != nil {
			return err
		}
		return output(result, []formatter.TableRow{result.Feed})
	},
}

var podcastsByTagCmd = &cobra.Command{
	Use:   "bytag",
	Short: "Get podcasts by namespace tag",
	RunE: func(cmd *cobra.Command, args []string) error {
		pv := mustBool(cmd, "podcast-value")
		pvts := mustBool(cmd, "podcast-value-time-split")
		if !pv && !pvts {
			return fmt.Errorf("one of --podcast-value or --podcast-value-time-split is required")
		}
		result, err := newClient().PodcastsByTag(cmd.Context(), client.PodcastsByTagOptions{
			PodcastValue:          pv,
			PodcastValueTimeSplit: pvts,
			Max:                   mustInt(cmd, "max"),
			StartAt:               mustInt(cmd, "start-at"),
		})
		if err != nil {
			return err
		}
		return output(result, toRows(result.Feeds))
	},
}

var podcastsByMediumCmd = &cobra.Command{
	Use:   "bymedium",
	Short: "Get podcasts by medium type",
	RunE: func(cmd *cobra.Command, args []string) error {
		medium := mustString(cmd, "medium")
		if medium == "" {
			return fmt.Errorf("--medium is required")
		}
		result, err := newClient().PodcastsByMedium(cmd.Context(), client.PodcastsByMediumOptions{
			Medium: medium,
			Max:    mustInt(cmd, "max"),
		})
		if err != nil {
			return err
		}
		return output(result, toRows(result.Feeds))
	},
}

var podcastsTrendingCmd = &cobra.Command{
	Use:   "trending",
	Short: "Get trending podcasts",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := newClient().PodcastsTrending(cmd.Context(), client.PodcastsTrendingOptions{
			Max:    mustInt(cmd, "max"),
			Since:  mustInt64(cmd, "since"),
			Lang:   mustString(cmd, "lang"),
			Cat:    mustString(cmd, "cat"),
			NotCat: mustString(cmd, "notcat"),
		})
		if err != nil {
			return err
		}
		return output(result, toRows(result.Feeds))
	},
}

var podcastsDeadCmd = &cobra.Command{
	Use:   "dead",
	Short: "Get dead feeds",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := newClient().PodcastsDead(cmd.Context())
		if err != nil {
			return err
		}
		return output(result, toRows(result.Feeds))
	},
}

var podcastsBatchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch get podcasts by GUIDs",
	RunE: func(cmd *cobra.Command, args []string) error {
		raw := mustString(cmd, "guids")
		if raw == "" {
			return fmt.Errorf("--guids is required")
		}
		guids := strings.Split(raw, ",")
		result, err := newClient().PodcastsBatchByGUID(cmd.Context(), guids)
		if err != nil {
			return err
		}
		return output(result, toRows(result.Feeds))
	},
}

func init() {
	rootCmd.AddCommand(podcastsCmd)

	podcastsByFeedIDCmd.Flags().Int("id", 0, "PodcastIndex Feed ID (required)")
	podcastsByFeedURLCmd.Flags().String("url", "", "Podcast feed URL (required)")
	podcastsByItunesIDCmd.Flags().Int("id", 0, "iTunes ID (required)")
	podcastsByGUIDCmd.Flags().String("guid", "", "Podcast GUID (required)")

	podcastsByTagCmd.Flags().Bool("podcast-value", false, "Feeds supporting podcast:value tag")
	podcastsByTagCmd.Flags().Bool("podcast-value-time-split", false, "Feeds supporting podcast:valueTimeSplit tag")
	podcastsByTagCmd.Flags().Int("max", 0, "Maximum number of results")
	podcastsByTagCmd.Flags().Int("start-at", 0, "Feed ID to start at for pagination")

	podcastsByMediumCmd.Flags().String("medium", "", "Medium type: audiobook, blog, film, music, newsletter, podcast, video (required)")
	podcastsByMediumCmd.Flags().Int("max", 0, "Maximum number of results")

	podcastsTrendingCmd.Flags().Int("max", 0, "Maximum number of results")
	podcastsTrendingCmd.Flags().Int64("since", 0, "Return items since this epoch timestamp")
	podcastsTrendingCmd.Flags().String("lang", "", "Language filter (comma-separated)")
	podcastsTrendingCmd.Flags().String("cat", "", "Category filter (comma-separated)")
	podcastsTrendingCmd.Flags().String("notcat", "", "Exclude categories (comma-separated)")

	podcastsBatchCmd.Flags().String("guids", "", "Comma-separated podcast GUIDs (required)")

	for _, c := range []*cobra.Command{
		podcastsByFeedIDCmd, podcastsByFeedURLCmd, podcastsByItunesIDCmd,
		podcastsByGUIDCmd, podcastsByTagCmd, podcastsByMediumCmd,
		podcastsTrendingCmd, podcastsDeadCmd, podcastsBatchCmd,
	} {
		podcastsCmd.AddCommand(c)
	}
}
