package media

import (
	"testing"
)

func TestDetectPlayer(t *testing.T) {
	player, err := DetectPlayer()
	if err != nil {
		t.Skipf("no media player found in PATH, skipping: %v", err)
	}
	if player == "" {
		t.Fatal("DetectPlayer returned empty string with nil error")
	}
}

func TestDetectPlayerCandidates(t *testing.T) {
	if len(defaultPlayers) != 3 {
		t.Fatalf("expected 3 default players, got %d", len(defaultPlayers))
	}
	expected := []string{"mpv", "vlc", "ffplay"}
	for i, name := range expected {
		if defaultPlayers[i] != name {
			t.Errorf("defaultPlayers[%d] = %q, want %q", i, defaultPlayers[i], name)
		}
	}
}
