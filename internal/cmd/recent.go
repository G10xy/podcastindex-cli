package cmd

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/G10xy/podcastindex-cli/internal/client"
	"github.com/G10xy/podcastindex-cli/internal/formatter"
	"github.com/G10xy/podcastindex-cli/pkg/models"
)

var recentCmd = &cobra.Command{
	Use:   "recent",
	Short: "Query recent data from the PodcastIndex",
}

var recentEpisodesCmd = &cobra.Command{
	Use:   "episodes",
	Short: "Most recent episodes globally",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := newClient().RecentEpisodes(cmd.Context(), client.RecentEpisodesOptions{
			Max:           mustInt(cmd, "max"),
			ExcludeString: mustString(cmd, "exclude"),
			Before:        mustInt(cmd, "before"),
			FullText:      mustBool(cmd, "fulltext"),
		})
		if err != nil {
			return err
		}
		return output(result, toRows(result.Items))
	},
}

var recentFeedsCmd = &cobra.Command{
	Use:   "feeds",
	Short: "Most recent feeds",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := newClient().RecentFeeds(cmd.Context(), client.RecentFeedsOptions{
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

var recentNewFeedsCmd = &cobra.Command{
	Use:   "newfeeds",
	Short: "New feeds in last 24 hours",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := newClient().RecentNewFeeds(cmd.Context(), client.RecentNewFeedsOptions{
			Max:    mustInt(cmd, "max"),
			Since:  mustInt64(cmd, "since"),
			FeedID: mustString(cmd, "feedid"),
			Desc:   mustBool(cmd, "desc"),
		})
		if err != nil {
			return err
		}
		return output(result, toRows(result.Feeds))
	},
}

var recentNewValueFeedsCmd = &cobra.Command{
	Use:   "newvaluefeeds",
	Short: "Feeds that recently added value tags",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := newClient().RecentNewValueFeeds(cmd.Context(), client.RecentNewValueFeedsOptions{
			Max:   mustInt(cmd, "max"),
			Since: mustInt64(cmd, "since"),
		})
		if err != nil {
			return err
		}
		return output(result, toRows(result.Feeds))
	},
}

var recentDataCmd = &cobra.Command{
	Use:   "data",
	Short: "Recent data (feeds + items)",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := newClient().RecentData(cmd.Context(), client.RecentDataOptions{
			Max:   mustInt(cmd, "max"),
			Since: mustInt64(cmd, "since"),
		})
		if err != nil {
			return err
		}
		// For JSON output the full response; for table/plain show feeds as rows
		rows := make([]formatter.TableRow, len(result.Data.Feeds))
		for i, f := range result.Data.Feeds {
			rows[i] = recentDataRow(f)
		}
		return output(result, rows)
	},
}

var recentSoundbitesCmd = &cobra.Command{
	Use:   "soundbites",
	Short: "Most recent soundbites",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := newClient().RecentSoundbites(cmd.Context(), mustInt(cmd, "max"))
		if err != nil {
			return err
		}
		return output(result, toRows(result.Items))
	},
}

func init() {
	rootCmd.AddCommand(recentCmd)

	recentEpisodesCmd.Flags().Int("max", 0, "Maximum number of results")
	recentEpisodesCmd.Flags().String("exclude", "", "Discard items containing this string")
	recentEpisodesCmd.Flags().Int("before", 0, "Episode ID to paginate before")
	recentEpisodesCmd.Flags().Bool("fulltext", false, "Return full text description")

	recentFeedsCmd.Flags().Int("max", 0, "Maximum number of results")
	recentFeedsCmd.Flags().Int64("since", 0, "Items since this epoch timestamp")
	recentFeedsCmd.Flags().String("lang", "", "Language filter")
	recentFeedsCmd.Flags().String("cat", "", "Category filter")
	recentFeedsCmd.Flags().String("notcat", "", "Exclude categories")

	recentNewFeedsCmd.Flags().Int("max", 0, "Maximum number of results")
	recentNewFeedsCmd.Flags().Int64("since", 0, "Items since this epoch timestamp")
	recentNewFeedsCmd.Flags().String("feedid", "", "Feed ID to start from")
	recentNewFeedsCmd.Flags().Bool("desc", false, "Display in descending order")

	for _, c := range []*cobra.Command{recentNewValueFeedsCmd, recentDataCmd} {
		c.Flags().Int("max", 0, "Maximum number of results")
		c.Flags().Int64("since", 0, "Items since this epoch timestamp")
	}

	recentSoundbitesCmd.Flags().Int("max", 0, "Maximum number of results")

	for _, c := range []*cobra.Command{
		recentEpisodesCmd, recentFeedsCmd, recentNewFeedsCmd,
		recentNewValueFeedsCmd, recentDataCmd, recentSoundbitesCmd,
	} {
		recentCmd.AddCommand(c)
	}
}

// recentDataRow wraps RecentDataItem to satisfy formatter.TableRow.
type recentDataRow models.RecentDataItem

func (r recentDataRow) TableHeaders() []string {
	return []string{"FEED ID", "FEED TITLE", "EPISODE TITLE"}
}

func (r recentDataRow) TableRow() []string {
	return []string{
		strconv.Itoa(r.FeedID),
		r.FeedTitle,
		r.EpisodeTitle,
	}
}
