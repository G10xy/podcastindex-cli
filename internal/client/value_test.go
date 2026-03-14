package client

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

const valueJSON = `{"status":"true","value":{"model":{"type":"lightning","method":"keysend"},"destinations":[{"name":"Host","address":"abc123","type":"node","split":90}]}}`

func TestValueByFeedID(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("id") != "920666" {
			t.Errorf("id = %q, want 920666", r.URL.Query().Get("id"))
		}
		if r.Header.Get("X-Auth-Key") != "" {
			t.Error("value endpoints should not send auth headers")
		}
		w.Write([]byte(valueJSON))
	}))
	defer srv.Close()

	c := NewClient("key", "secret", WithBaseURL(srv.URL))
	result, err := c.ValueByFeedID(context.Background(), 920666)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Value.Destinations) != 1 {
		t.Fatalf("got %d destinations, want 1", len(result.Value.Destinations))
	}
	if result.Value.Destinations[0].Split != 90 {
		t.Errorf("split = %d, want 90", result.Value.Destinations[0].Split)
	}
}

func TestValueByFeedURL(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("url") == "" {
			t.Error("expected url param")
		}
		w.Write([]byte(valueJSON))
	}))
	defer srv.Close()

	c := NewClient("key", "secret", WithBaseURL(srv.URL))
	result, err := c.ValueByFeedURL(context.Background(), "https://example.com/feed.xml")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Value.Model.Type != "lightning" {
		t.Errorf("model type = %q", result.Value.Model.Type)
	}
}

func TestValueByEpisodeGUID(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("podcastguid") != "pg1" {
			t.Errorf("podcastguid = %q", r.URL.Query().Get("podcastguid"))
		}
		if r.URL.Query().Get("episodeguid") != "eg1" {
			t.Errorf("episodeguid = %q", r.URL.Query().Get("episodeguid"))
		}
		w.Write([]byte(valueJSON))
	}))
	defer srv.Close()

	c := NewClient("key", "secret", WithBaseURL(srv.URL))
	_, err := c.ValueByEpisodeGUID(context.Background(), "pg1", "eg1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValueBatchByEpisodeGUID(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %q, want POST", r.Method)
		}
		body, _ := io.ReadAll(r.Body)
		if len(body) == 0 {
			t.Error("expected non-empty body")
		}
		w.Write([]byte(`{"status":"true","items":{"pg1":{"eg1":{"model":{"type":"lightning"}}}}}`))
	}))
	defer srv.Close()

	c := NewClient("key", "secret", WithBaseURL(srv.URL))
	result, err := c.ValueBatchByEpisodeGUID(context.Background(), map[string][]string{"pg1": {"eg1"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Items == nil {
		t.Error("expected non-nil items")
	}
}
