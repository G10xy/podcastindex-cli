package media

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestDownload(t *testing.T) {
	content := "fake audio content for testing"
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", "30")
		w.Write([]byte(content))
	}))
	defer srv.Close()

	dir := t.TempDir()

	dest, err := Download(context.Background(), srv.Client(), DownloadOptions{
		URL:      srv.URL + "/episode.mp3",
		Dir:      dir,
		Filename: "test.mp3",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := filepath.Join(dir, "test.mp3")
	if dest != expected {
		t.Fatalf("dest = %q, want %q", dest, expected)
	}

	data, err := os.ReadFile(dest)
	if err != nil {
		t.Fatalf("reading file: %v", err)
	}
	if string(data) != content {
		t.Fatalf("file content = %q, want %q", string(data), content)
	}
}

func TestDownloadDeriveFilename(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("data"))
	}))
	defer srv.Close()

	dir := t.TempDir()

	dest, err := Download(context.Background(), srv.Client(), DownloadOptions{
		URL: srv.URL + "/path/to/my-episode.mp3",
		Dir: dir,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if filepath.Base(dest) != "my-episode.mp3" {
		t.Fatalf("derived filename = %q, want %q", filepath.Base(dest), "my-episode.mp3")
	}
}

func TestDownloadWithProgress(t *testing.T) {
	content := "some audio data"
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(content))
	}))
	defer srv.Close()

	dir := t.TempDir()

	dest, err := Download(context.Background(), srv.Client(), DownloadOptions{
		URL:      srv.URL + "/ep.mp3",
		Dir:      dir,
		Filename: "ep.mp3",
		Output:   os.Stderr,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(dest)
	if err != nil {
		t.Fatalf("reading file: %v", err)
	}
	if string(data) != content {
		t.Fatalf("file content mismatch")
	}
}

func TestDownloadHTTPError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()

	dir := t.TempDir()

	_, err := Download(context.Background(), srv.Client(), DownloadOptions{
		URL: srv.URL + "/missing.mp3",
		Dir: dir,
	})
	if err == nil {
		t.Fatal("expected error for 404 response")
	}
}

func TestDeriveFilename(t *testing.T) {
	tests := []struct {
		url  string
		want string
	}{
		{"https://example.com/podcast/episode-42.mp3", "episode-42.mp3"},
		{"https://cdn.example.com/path/to/audio.ogg?token=abc", "audio.ogg"},
		{"https://example.com/", ""},   // falls back to generated name
	}
	for _, tt := range tests {
		got := deriveFilename(tt.url)
		if tt.want == "" {
			// Just check it's not empty (will be a generated name)
			if got == "" {
				t.Errorf("deriveFilename(%q) returned empty", tt.url)
			}
		} else if got != tt.want {
			t.Errorf("deriveFilename(%q) = %q, want %q", tt.url, got, tt.want)
		}
	}
}

func TestSanitizeFilename(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"normal.mp3", "normal.mp3"},
		{"file/with/slashes.mp3", "file_with_slashes.mp3"},
		{"file:with:colons.mp3", "file_with_colons.mp3"},
		{" spaces.mp3 ", "spaces.mp3"},
	}
	for _, tt := range tests {
		got := sanitizeFilename(tt.input)
		if got != tt.want {
			t.Errorf("sanitizeFilename(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}
