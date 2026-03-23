package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"

	"github.com/G10xy/podcastindex-cli/internal/client"
	"github.com/G10xy/podcastindex-cli/internal/media"
	"github.com/G10xy/podcastindex-cli/pkg/models"
)

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download an episode audio file",
	Long: `Download a single episode by ID, GUID, or direct URL.
Use --all with --feedid, --feedurl, or --podcastguid to download all episodes from a feed.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		all := mustBool(cmd, "all")
		if all {
			return runBatchDownload(cmd)
		}
		return runSingleDownload(cmd)
	},
}

func init() {
	// Single episode flags
	downloadCmd.Flags().Int("id", 0, "Episode ID")
	downloadCmd.Flags().String("guid", "", "Episode GUID")
	downloadCmd.Flags().String("url", "", "Direct enclosure URL (skips API lookup)")

	// Batch flags
	downloadCmd.Flags().Bool("all", false, "Download all episodes from a feed")
	downloadCmd.Flags().Int("feedid", 0, "Feed ID (used with --all)")
	downloadCmd.Flags().String("feedurl", "", "Feed URL (used with --all)")
	downloadCmd.Flags().String("podcastguid", "", "Podcast GUID (used with --all)")
	downloadCmd.Flags().Int("workers", 3, "Number of concurrent downloads (used with --all)")

	// Output flags
	downloadCmd.Flags().String("dir", ".", "Output directory")
	downloadCmd.Flags().String("filename", "", "Override output filename (single episode only)")

	rootCmd.AddCommand(downloadCmd)
}

func runSingleDownload(cmd *cobra.Command) error {
	id := mustInt(cmd, "id")
	guid := mustString(cmd, "guid")
	directURL := mustString(cmd, "url")

	enclosureURL, title, err := resolveEpisode(cmd, id, guid, directURL)
	if err != nil {
		return err
	}

	if title != "" {
		fmt.Fprintf(os.Stderr, "Downloading %q...\n", title)
	}

	dest, err := media.Download(cmd.Context(), http.DefaultClient, media.DownloadOptions{
		URL:      enclosureURL,
		Dir:      mustString(cmd, "dir"),
		Filename: mustString(cmd, "filename"),
		Output:   os.Stderr,
	})
	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Saved to %s\n", dest)
	fmt.Println(dest)
	return nil
}

func runBatchDownload(cmd *cobra.Command) error {
	feedID := mustInt(cmd, "feedid")
	feedURL := mustString(cmd, "feedurl")
	podcastGUID := mustString(cmd, "podcastguid")

	if feedID == 0 && feedURL == "" && podcastGUID == "" {
		return fmt.Errorf("--all requires one of --feedid, --feedurl, or --podcastguid")
	}

	c := newClient()
	ctx := cmd.Context()

	episodes, err := fetchAllEpisodes(ctx, c, feedID, feedURL, podcastGUID)
	if err != nil {
		return err
	}

	if len(episodes) == 0 {
		fmt.Fprintln(os.Stderr, "No episodes found.")
		return nil
	}

	dir := mustString(cmd, "dir")
	workers := mustInt(cmd, "workers")

	fmt.Fprintf(os.Stderr, "Downloading %d episodes to %s (workers: %d)...\n", len(episodes), dir, workers)

	result, err := media.BatchDownload(ctx, http.DefaultClient, media.BatchDownloadOptions{
		Episodes: episodes,
		Dir:      dir,
		Workers:  workers,
		Output:   os.Stderr,
	})
	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "\nDone. Downloaded %d episodes (%s total)",
		result.TotalFiles, formatBytesForSummary(result.TotalBytes))
	if len(result.Errors) > 0 {
		fmt.Fprintf(os.Stderr, ", %d failed", len(result.Errors))
	}
	fmt.Fprintln(os.Stderr)

	return nil
}

func fetchAllEpisodes(ctx context.Context, c *client.Client, feedID int, feedURL, podcastGUID string) ([]models.Episode, error) {
	switch {
	case feedID != 0:
		result, err := c.EpisodesByFeedID(ctx, client.EpisodesByFeedIDOptions{
			ID:  fmt.Sprintf("%d", feedID),
			Max: 1000,
		})
		if err != nil {
			return nil, err
		}
		return result.Items, nil

	case feedURL != "":
		result, err := c.EpisodesByFeedURL(ctx, client.EpisodesByFeedURLOptions{
			URL: feedURL,
			Max: 1000,
		})
		if err != nil {
			return nil, err
		}
		return result.Items, nil

	case podcastGUID != "":
		result, err := c.EpisodesByPodcastGUID(ctx, client.EpisodesByPodcastGUIDOptions{
			GUID: podcastGUID,
			Max:  1000,
		})
		if err != nil {
			return nil, err
		}
		return result.Items, nil
	}

	return nil, fmt.Errorf("one of --feedid, --feedurl, or --podcastguid is required")
}

func formatBytesForSummary(b int64) string {
	const (
		kb = 1024
		mb = kb * 1024
		gb = mb * 1024
	)
	switch {
	case b >= gb:
		return fmt.Sprintf("%.1f GB", float64(b)/float64(gb))
	case b >= mb:
		return fmt.Sprintf("%.1f MB", float64(b)/float64(mb))
	case b >= kb:
		return fmt.Sprintf("%.1f KB", float64(b)/float64(kb))
	default:
		return fmt.Sprintf("%d B", b)
	}
}
