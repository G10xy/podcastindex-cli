package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/G10xy/podcastindex-cli/internal/client"
	"github.com/G10xy/podcastindex-cli/internal/media"
)

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Stream an episode to a local media player",
	Long:  "Play an episode by resolving its audio URL and launching a local media player (mpv, vlc, or ffplay).",
	RunE: func(cmd *cobra.Command, args []string) error {
		id := mustInt(cmd, "id")
		guid := mustString(cmd, "guid")
		directURL := mustString(cmd, "url")

		enclosureURL, title, err := resolveEpisode(cmd, id, guid, directURL)
		if err != nil {
			return err
		}

		player := mustString(cmd, "player")
		if player == "" {
			player = viper.GetString("player")
		}
		if player == "" {
			player, err = media.DetectPlayer()
			if err != nil {
				return err
			}
		}

		if title != "" {
			fmt.Fprintf(os.Stderr, "Playing %q with %s...\n", title, player)
		} else {
			fmt.Fprintf(os.Stderr, "Playing with %s...\n", player)
		}

		return media.Play(cmd.Context(), player, enclosureURL)
	},
}

func init() {
	playCmd.Flags().Int("id", 0, "Episode ID")
	playCmd.Flags().String("guid", "", "Episode GUID")
	playCmd.Flags().String("url", "", "Direct enclosure URL (skips API lookup)")
	playCmd.Flags().String("player", "", "Media player to use (default: auto-detect mpv/vlc/ffplay)")
	rootCmd.AddCommand(playCmd)
}

// resolveEpisode returns the enclosure URL and title for an episode
// identified by ID, GUID, or direct URL.
func resolveEpisode(cmd *cobra.Command, id int, guid, directURL string) (enclosureURL, title string, err error) {
	if directURL != "" {
		return directURL, "", nil
	}

	if id == 0 && guid == "" {
		return "", "", fmt.Errorf("one of --id, --guid, or --url is required")
	}

	c := newClient()

	if id != 0 {
		result, err := c.EpisodeByID(cmd.Context(), id, false)
		if err != nil {
			return "", "", err
		}
		if result.Episode.EnclosureURL == "" {
			return "", "", fmt.Errorf("episode %d has no enclosure URL", id)
		}
		return result.Episode.EnclosureURL, result.Episode.Title, nil
	}

	result, err := c.EpisodeByGUID(cmd.Context(), client.EpisodeByGUIDOptions{GUID: guid})
	if err != nil {
		return "", "", err
	}
	if result.Episode.EnclosureURL == "" {
		return "", "", fmt.Errorf("episode %q has no enclosure URL", guid)
	}
	return result.Episode.EnclosureURL, result.Episode.Title, nil
}
