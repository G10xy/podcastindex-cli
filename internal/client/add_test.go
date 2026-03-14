package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddByFeedURL(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/add/byfeedurl" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.URL.Query().Get("url") != "https://example.com/feed.xml" {
			t.Errorf("url = %q", r.URL.Query().Get("url"))
		}
		if r.URL.Query().Get("chash") != "abc123" {
			t.Errorf("chash = %q, want abc123", r.URL.Query().Get("chash"))
		}
		w.Write([]byte(`{"status":"true","description":"Feed added","feedId":12345}`))
	}))
	defer srv.Close()

	c := NewClient("key", "secret", WithBaseURL(srv.URL))
	result, err := c.AddByFeedURL(context.Background(), AddByFeedURLOptions{
		URL:   "https://example.com/feed.xml",
		CHash: "abc123",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.FeedID != 12345 {
		t.Errorf("feedId = %d, want 12345", result.FeedID)
	}
}

func TestAddByItunesID(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("id") != "1441923632" {
			t.Errorf("id = %q, want 1441923632", r.URL.Query().Get("id"))
		}
		w.Write([]byte(`{"status":"true","description":"Feed added","feedId":67890}`))
	}))
	defer srv.Close()

	c := NewClient("key", "secret", WithBaseURL(srv.URL))
	result, err := c.AddByItunesID(context.Background(), 1441923632)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.FeedID != 67890 {
		t.Errorf("feedId = %d, want 67890", result.FeedID)
	}
}
