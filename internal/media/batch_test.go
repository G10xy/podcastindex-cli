package media

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/G10xy/podcastindex-cli/pkg/models"
)

func TestBatchDownload(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("audio data"))
	}))
	defer srv.Close()

	episodes := []models.Episode{
		{Title: "Episode 1", EnclosureURL: srv.URL + "/ep1.mp3"},
		{Title: "Episode 2", EnclosureURL: srv.URL + "/ep2.mp3"},
		{Title: "Episode 3", EnclosureURL: srv.URL + "/ep3.mp3"},
	}

	dir := t.TempDir()
	var output bytes.Buffer

	result, err := BatchDownload(context.Background(), srv.Client(), BatchDownloadOptions{
		Episodes: episodes,
		Dir:      dir,
		Workers:  2,
		Output:   &output,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.TotalFiles != 3 {
		t.Errorf("TotalFiles = %d, want 3", result.TotalFiles)
	}
	if len(result.Errors) != 0 {
		t.Errorf("got %d errors, want 0", len(result.Errors))
	}

	// Verify output contains completion messages
	out := output.String()
	if !strings.Contains(out, "Episode 1") || !strings.Contains(out, "Episode 2") || !strings.Contains(out, "Episode 3") {
		t.Errorf("output missing episode names: %q", out)
	}
}

func TestBatchDownloadSkipsEmptyURL(t *testing.T) {
	episodes := []models.Episode{
		{Title: "No URL Episode", EnclosureURL: ""},
	}

	dir := t.TempDir()
	var output bytes.Buffer

	result, err := BatchDownload(context.Background(), http.DefaultClient, BatchDownloadOptions{
		Episodes: episodes,
		Dir:      dir,
		Workers:  1,
		Output:   &output,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.TotalFiles != 0 {
		t.Errorf("TotalFiles = %d, want 0", result.TotalFiles)
	}
	if !strings.Contains(output.String(), "Skipped") {
		t.Errorf("expected 'Skipped' in output, got %q", output.String())
	}
}

func TestBatchDownloadEmpty(t *testing.T) {
	result, err := BatchDownload(context.Background(), http.DefaultClient, BatchDownloadOptions{
		Episodes: nil,
		Dir:      t.TempDir(),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.TotalFiles != 0 {
		t.Errorf("TotalFiles = %d, want 0", result.TotalFiles)
	}
}

func TestBatchDownloadConcurrencyLimit(t *testing.T) {
	var maxConcurrent atomic.Int32
	var current atomic.Int32

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := current.Add(1)
		for {
			m := maxConcurrent.Load()
			if c > m {
				if maxConcurrent.CompareAndSwap(m, c) {
					break
				}
			} else {
				break
			}
		}
		w.Write([]byte("data"))
		current.Add(-1)
	}))
	defer srv.Close()

	episodes := make([]models.Episode, 10)
	for i := range episodes {
		episodes[i] = models.Episode{
			Title:        "Ep",
			EnclosureURL: srv.URL + "/ep.mp3",
		}
	}

	dir := t.TempDir()
	_, err := BatchDownload(context.Background(), srv.Client(), BatchDownloadOptions{
		Episodes: episodes,
		Dir:      dir,
		Workers:  2,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if maxConcurrent.Load() > 2 {
		t.Errorf("max concurrent = %d, expected at most 2", maxConcurrent.Load())
	}
}
