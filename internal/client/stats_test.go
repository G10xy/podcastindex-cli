package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatsCurrent(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/stats/current" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Write([]byte(`{"status":"true","stats":{"feedCountTotal":4361712,"episodeCountTotal":134573938,"feedsWithNewEpisodes":74583,"feedsWithValueBlocks":11200}}`))
	}))
	defer srv.Close()

	c := NewClient("key", "secret", WithBaseURL(srv.URL))
	result, err := c.StatsCurrent(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Stats.FeedCountTotal != 4361712 {
		t.Errorf("feedCountTotal = %d", result.Stats.FeedCountTotal)
	}
}
