package media

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

var defaultPlayers = []string{"mpv", "vlc", "ffplay"}

// DetectPlayer searches PATH for a supported media player
// and returns the first found. Checks mpv, vlc, ffplay in order.
func DetectPlayer() (string, error) {
	for _, name := range defaultPlayers {
		if path, err := exec.LookPath(name); err == nil {
			return path, nil
		}
	}
	return "", fmt.Errorf("no supported media player found in PATH (tried: mpv, vlc, ffplay)")
}

// Play launches the given player binary with the enclosure URL
// and blocks until the player exits.
func Play(ctx context.Context, player, url string) error {
	cmd := exec.CommandContext(ctx, player, url)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("player %q exited with error: %w", player, err)
	}
	return nil
}
