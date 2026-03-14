package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecentEpisodes(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/recent/episodes" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.URL.Query().Get("max") != "5" {
			t.Errorf("max = %q, want 5", r.URL.Query().Get("max"))
		}
		if r.URL.Query().Get("excludeString") != "religion" {
			t.Errorf("excludeString = %q", r.URL.Query().Get("excludeString"))
		}
		w.Write([]byte(`{"status":"true","items":[{"id":1,"title":"Recent Ep"}],"count":1}`))
	}))
	defer srv.Close()

	c := NewClient("key", "secret", WithBaseURL(srv.URL))
	result, err := c.RecentEpisodes(context.Background(), RecentEpisodesOptions{Max: 5, ExcludeString: "religion"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Items) != 1 {
		t.Fatalf("got %d items, want 1", len(result.Items))
	}
}

func TestRecentFeeds(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("lang") != "en" {
			t.Errorf("lang = %q", r.URL.Query().Get("lang"))
		}
		w.Write([]byte(`{"status":"true","feeds":[{"id":1,"title":"Recent Feed"}],"count":1}`))
	}))
	defer srv.Close()

	c := NewClient("key", "secret", WithBaseURL(srv.URL))
	result, err := c.RecentFeeds(context.Background(), RecentFeedsOptions{Lang: "en"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Feeds[0].Title != "Recent Feed" {
		t.Errorf("title = %q", result.Feeds[0].Title)
	}
}

func TestRecentNewFeeds(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, ok := r.URL.Query()["desc"]; !ok {
			t.Error("expected desc param")
		}
		w.Write([]byte(`{"status":"true","feeds":[{"id":100,"url":"https://example.com/feed.xml","status":"confirmed"}],"count":1}`))
	}))
	defer srv.Close()

	c := NewClient("key", "secret", WithBaseURL(srv.URL))
	result, err := c.RecentNewFeeds(context.Background(), RecentNewFeedsOptions{Desc: true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Feeds[0].ID != 100 {
		t.Errorf("id = %d, want 100", result.Feeds[0].ID)
	}
}

func TestRecentNewValueFeeds(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/recent/newvaluefeeds" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Write([]byte(`{"status":"true","feeds":[{"id":50,"title":"Value Feed"}],"count":1}`))
	}))
	defer srv.Close()

	c := NewClient("key", "secret", WithBaseURL(srv.URL))
	result, err := c.RecentNewValueFeeds(context.Background(), RecentNewValueFeedsOptions{Max: 10})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Feeds) != 1 {
		t.Fatalf("got %d feeds, want 1", len(result.Feeds))
	}
}

func TestRecentData(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status":"true","data":{"feeds":[{"feedId":1,"feedTitle":"F1"}],"items":[{"episodeId":10,"episodeTitle":"E1"}]}}`))
	}))
	defer srv.Close()

	c := NewClient("key", "secret", WithBaseURL(srv.URL))
	result, err := c.RecentData(context.Background(), RecentDataOptions{Max: 100})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Data.Feeds) != 1 {
		t.Fatalf("got %d feeds, want 1", len(result.Data.Feeds))
	}
	if result.Data.Items[0].EpisodeID != 10 {
		t.Errorf("episodeId = %d, want 10", result.Data.Items[0].EpisodeID)
	}
}

func TestRecentSoundbites(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/recent/soundbites" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Write([]byte(`{"status":"true","items":[{"episodeId":200,"title":"Bite","startTime":30,"duration":60}],"count":1}`))
	}))
	defer srv.Close()

	c := NewClient("key", "secret", WithBaseURL(srv.URL))
	result, err := c.RecentSoundbites(context.Background(), 20)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Items[0].Duration != 60 {
		t.Errorf("duration = %d, want 60", result.Items[0].Duration)
	}
}
