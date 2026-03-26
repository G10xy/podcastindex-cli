package cmd

import (
	"github.com/spf13/cobra"

	"github.com/G10xy/podcastindex-cli/internal/client"
	"github.com/G10xy/podcastindex-cli/internal/formatter"
)

var episodesCmd = &cobra.Command{
	Use:   "episodes",
	Short: "Query episodes from the PodcastIndex",
}

var episodesByFeedIDCmd = &cobra.Command{
	Use:   "byfeedid",
	Short: "Get episodes by Feed ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := newClient().EpisodesByFeedID(cmd.Context(), client.EpisodesByFeedIDOptions{
			ID:        mustString(cmd, "id"),
			Since:     mustInt64(cmd, "since"),
			Max:       mustInt(cmd, "max"),
			Enclosure: mustString(cmd, "enclosure"),
			FullText:  mustBool(cmd, "fulltext"),
			Newest:    mustBool(cmd, "newest"),
		})
		if err != nil {
			return err
		}
		return output(result, toRows(result.Items))
	},
}

var episodesByFeedURLCmd = &cobra.Command{
	Use:   "byfeedurl",
	Short: "Get episodes by feed URL",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := newClient().EpisodesByFeedURL(cmd.Context(), client.EpisodesByFeedURLOptions{
			URL:      mustString(cmd, "url"),
			Since:    mustInt64(cmd, "since"),
			Max:      mustInt(cmd, "max"),
			FullText: mustBool(cmd, "fulltext"),
		})
		if err != nil {
			return err
		}
		return output(result, toRows(result.Items))
	},
}

var episodesByPodcastGUIDCmd = &cobra.Command{
	Use:   "bypodcastguid",
	Short: "Get episodes by Podcast GUID",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := newClient().EpisodesByPodcastGUID(cmd.Context(), client.EpisodesByPodcastGUIDOptions{
			GUID:     mustString(cmd, "guid"),
			Since:    mustInt64(cmd, "since"),
			Max:      mustInt(cmd, "max"),
			FullText: mustBool(cmd, "fulltext"),
		})
		if err != nil {
			return err
		}
		return output(result, toRows(result.Items))
	},
}

var episodesByItunesIDCmd = &cobra.Command{
	Use:   "byitunesid",
	Short: "Get episodes by iTunes ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := newClient().EpisodesByItunesID(cmd.Context(), client.EpisodesByItunesIDOptions{
			ID:        mustInt(cmd, "id"),
			Since:     mustInt64(cmd, "since"),
			Max:       mustInt(cmd, "max"),
			Enclosure: mustString(cmd, "enclosure"),
			FullText:  mustBool(cmd, "fulltext"),
		})
		if err != nil {
			return err
		}
		return output(result, toRows(result.Items))
	},
}

var episodesByIDCmd = &cobra.Command{
	Use:   "byid",
	Short: "Get single episode by ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := newClient().EpisodeByID(cmd.Context(), mustInt(cmd, "id"), mustBool(cmd, "fulltext"))
		if err != nil {
			return err
		}
		return output(result, []formatter.TableRow{result.Episode})
	},
}

var episodesByGUIDCmd = &cobra.Command{
	Use:   "byguid",
	Short: "Get single episode by GUID",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := newClient().EpisodeByGUID(cmd.Context(), client.EpisodeByGUIDOptions{
			GUID:        mustString(cmd, "guid"),
			FeedURL:     mustString(cmd, "feedurl"),
			FeedID:      mustString(cmd, "feedid"),
			PodcastGUID: mustString(cmd, "podcastguid"),
			FullText:    mustBool(cmd, "fulltext"),
		})
		if err != nil {
			return err
		}
		return output(result, []formatter.TableRow{result.Episode})
	},
}

var episodesLiveCmd = &cobra.Command{
	Use:   "live",
	Short: "Get live episodes",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := newClient().EpisodesLive(cmd.Context(), mustInt(cmd, "max"))
		if err != nil {
			return err
		}
		return output(result, toRows(result.Items))
	},
}

var episodesRandomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get random episodes",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := newClient().EpisodesRandom(cmd.Context(), client.EpisodesRandomOptions{
			Max:      mustInt(cmd, "max"),
			Lang:     mustString(cmd, "lang"),
			Cat:      mustString(cmd, "cat"),
			NotCat:   mustString(cmd, "notcat"),
			FullText: mustBool(cmd, "fulltext"),
		})
		if err != nil {
			return err
		}
		return output(result, toRows(result.Episodes))
	},
}

func init() {
	rootCmd.AddCommand(episodesCmd)

	// Common flags for multi-episode commands
	for _, c := range []*cobra.Command{episodesByFeedIDCmd, episodesByFeedURLCmd, episodesByPodcastGUIDCmd, episodesByItunesIDCmd} {
		c.Flags().Int64("since", 0, "Return items since this epoch timestamp")
		c.Flags().Int("max", 0, "Maximum number of results")
		c.Flags().Bool("fulltext", false, "Return full text description")
	}

	episodesByFeedIDCmd.Flags().String("id", "", "Feed ID or comma-separated IDs (required)")
	episodesByFeedIDCmd.MarkFlagRequired("id")
	episodesByFeedIDCmd.Flags().String("enclosure", "", "Filter by enclosure URL")
	episodesByFeedIDCmd.Flags().Bool("newest", false, "Only most recent episode per feed")

	episodesByFeedURLCmd.Flags().String("url", "", "Feed URL (required)")
	episodesByFeedURLCmd.MarkFlagRequired("url")

	episodesByPodcastGUIDCmd.Flags().String("guid", "", "Podcast GUID (required)")
	episodesByPodcastGUIDCmd.MarkFlagRequired("guid")

	episodesByItunesIDCmd.Flags().Int("id", 0, "iTunes Feed ID (required)")
	episodesByItunesIDCmd.MarkFlagRequired("id")
	episodesByItunesIDCmd.Flags().String("enclosure", "", "Filter by enclosure URL")

	episodesByIDCmd.Flags().Int("id", 0, "Episode ID (required)")
	episodesByIDCmd.MarkFlagRequired("id")
	episodesByIDCmd.Flags().Bool("fulltext", false, "Return full text description")

	episodesByGUIDCmd.Flags().String("guid", "", "Episode GUID (required)")
	episodesByGUIDCmd.MarkFlagRequired("guid")
	episodesByGUIDCmd.Flags().String("feedurl", "", "Feed URL")
	episodesByGUIDCmd.Flags().String("feedid", "", "Feed ID")
	episodesByGUIDCmd.Flags().String("podcastguid", "", "Podcast GUID")
	episodesByGUIDCmd.Flags().Bool("fulltext", false, "Return full text description")

	episodesLiveCmd.Flags().Int("max", 0, "Maximum number of results")

	episodesRandomCmd.Flags().Int("max", 0, "Maximum number of results")
	episodesRandomCmd.Flags().String("lang", "", "Language filter (comma-separated)")
	episodesRandomCmd.Flags().String("cat", "", "Category filter")
	episodesRandomCmd.Flags().String("notcat", "", "Exclude categories")
	episodesRandomCmd.Flags().Bool("fulltext", false, "Return full text description")

	for _, c := range []*cobra.Command{
		episodesByFeedIDCmd, episodesByFeedURLCmd, episodesByPodcastGUIDCmd,
		episodesByItunesIDCmd, episodesByIDCmd, episodesByGUIDCmd,
		episodesLiveCmd, episodesRandomCmd,
	} {
		episodesCmd.AddCommand(c)
	}
}
