package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

const episodesListJSON = `{"status":"true","items":[{"id":1001,"title":"Ep 1","feedTitle":"Pod A"},{"id":1002,"title":"Ep 2","feedTitle":"Pod A"}],"count":2}`
const singleEpisodeJSON = `{"status":"true","episode":{"id":1001,"title":"Ep 1","feedTitle":"Pod A"}}`

func TestEpisodesByFeedID(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/episodes/byfeedid" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.URL.Query().Get("id") != "75075" {
			t.Errorf("id = %q, want 75075", r.URL.Query().Get("id"))
		}
		if _, ok := r.URL.Query()["newest"]; !ok {
			t.Error("expected newest param")
		}
		w.Write([]byte(episodesListJSON))
	}))
	defer srv.Close()

	c := NewClient("key", "secret", WithBaseURL(srv.URL))
	result, err := c.EpisodesByFeedID(context.Background(), EpisodesByFeedIDOptions{ID: "75075", Newest: true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Items) != 2 {
		t.Fatalf("got %d items, want 2", len(result.Items))
	}
}

func TestEpisodesByFeedURL(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("url") != "https://example.com/feed.xml" {
			t.Errorf("url = %q", r.URL.Query().Get("url"))
		}
		w.Write([]byte(episodesListJSON))
	}))
	defer srv.Close()

	c := NewClient("key", "secret", WithBaseURL(srv.URL))
	result, err := c.EpisodesByFeedURL(context.Background(), EpisodesByFeedURLOptions{URL: "https://example.com/feed.xml"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Items) != 2 {
		t.Fatalf("got %d items, want 2", len(result.Items))
	}
}

func TestEpisodeByID(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("id") != "1001" {
			t.Errorf("id = %q, want 1001", r.URL.Query().Get("id"))
		}
		w.Write([]byte(singleEpisodeJSON))
	}))
	defer srv.Close()

	c := NewClient("key", "secret", WithBaseURL(srv.URL))
	result, err := c.EpisodeByID(context.Background(), 1001, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Episode.Title != "Ep 1" {
		t.Errorf("title = %q, want Ep 1", result.Episode.Title)
	}
}

func TestEpisodeByGUID(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("guid") != "PC2084" {
			t.Errorf("guid = %q", r.URL.Query().Get("guid"))
		}
		if r.URL.Query().Get("feedid") != "920666" {
			t.Errorf("feedid = %q", r.URL.Query().Get("feedid"))
		}
		w.Write([]byte(singleEpisodeJSON))
	}))
	defer srv.Close()

	c := NewClient("key", "secret", WithBaseURL(srv.URL))
	result, err := c.EpisodeByGUID(context.Background(), EpisodeByGUIDOptions{GUID: "PC2084", FeedID: "920666"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Episode.ID != 1001 {
		t.Errorf("id = %d, want 1001", result.Episode.ID)
	}
}

func TestEpisodesLive(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/episodes/live" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Write([]byte(`{"status":"true","items":[{"id":5000,"title":"Live Now"}],"count":1}`))
	}))
	defer srv.Close()

	c := NewClient("key", "secret", WithBaseURL(srv.URL))
	result, err := c.EpisodesLive(context.Background(), 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Items[0].Title != "Live Now" {
		t.Errorf("title = %q", result.Items[0].Title)
	}
}

func TestEpisodesRandom(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("lang") != "en" {
			t.Errorf("lang = %q, want en", r.URL.Query().Get("lang"))
		}
		w.Write([]byte(`{"status":"true","episodes":[{"id":9999,"title":"Random Ep"}],"count":1}`))
	}))
	defer srv.Close()

	c := NewClient("key", "secret", WithBaseURL(srv.URL))
	result, err := c.EpisodesRandom(context.Background(), EpisodesRandomOptions{Max: 1, Lang: "en"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Episodes[0].ID != 9999 {
		t.Errorf("id = %d, want 9999", result.Episodes[0].ID)
	}
}
