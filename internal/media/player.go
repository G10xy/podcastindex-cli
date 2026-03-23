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

// ResolvePlayer validates a player name by looking it up in PATH.
// Returns the absolute path to the player binary.
func ResolvePlayer(name string) (string, error) {
	path, err := exec.LookPath(name)
	if err != nil {
		return "", fmt.Errorf("player %q not found in PATH: %w", name, err)
	}
	return path, nil
}

// Play launches the given player binary with the enclosure URL
// and blocks until the player exits. The player must be an absolute
// path or resolvable via PATH.
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
