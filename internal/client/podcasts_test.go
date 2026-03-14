package client

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPodcastByFeedID(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/podcasts/byfeedid" {
			t.Errorf("path = %q, want /podcasts/byfeedid", r.URL.Path)
		}
		if r.URL.Query().Get("id") != "75075" {
			t.Errorf("id = %q, want 75075", r.URL.Query().Get("id"))
		}
		w.Write([]byte(`{"status":"true","feed":{"id":75075,"title":"Batman University"}}`))
	}))
	defer srv.Close()

	c := NewClient("key", "secret", WithBaseURL(srv.URL))
	result, err := c.PodcastByFeedID(context.Background(), 75075)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Feed.Title != "Batman University" {
		t.Errorf("title = %q, want Batman University", result.Feed.Title)
	}
}

func TestPodcastByFeedURL(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("url") != "https://example.com/feed.xml" {
			t.Errorf("url = %q", r.URL.Query().Get("url"))
		}
		w.Write([]byte(`{"status":"true","feed":{"id":1,"title":"Test Feed"}}`))
	}))
	defer srv.Close()

	c := NewClient("key", "secret", WithBaseURL(srv.URL))
	result, err := c.PodcastByFeedURL(context.Background(), "https://example.com/feed.xml")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Feed.ID != 1 {
		t.Errorf("id = %d, want 1", result.Feed.ID)
	}
}

func TestPodcastsByTag(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, ok := r.URL.Query()["podcast-value"]; !ok {
			t.Error("expected podcast-value param")
		}
		if r.URL.Query().Get("max") != "5" {
			t.Errorf("max = %q, want 5", r.URL.Query().Get("max"))
		}
		w.Write([]byte(`{"status":"true","feeds":[{"id":1,"title":"Value Feed"}],"count":1}`))
	}))
	defer srv.Close()

	c := NewClient("key", "secret", WithBaseURL(srv.URL))
	result, err := c.PodcastsByTag(context.Background(), PodcastsByTagOptions{
		PodcastValue: true,
		Max:          5,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Feeds) != 1 {
		t.Fatalf("got %d feeds, want 1", len(result.Feeds))
	}
}

func TestPodcastsTrending(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("lang") != "en" {
			t.Errorf("lang = %q, want en", r.URL.Query().Get("lang"))
		}
		w.Write([]byte(`{"status":"true","feeds":[{"id":10,"title":"Trending Pod"}],"count":1}`))
	}))
	defer srv.Close()

	c := NewClient("key", "secret", WithBaseURL(srv.URL))
	result, err := c.PodcastsTrending(context.Background(), PodcastsTrendingOptions{Lang: "en"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Feeds[0].Title != "Trending Pod" {
		t.Errorf("title = %q, want Trending Pod", result.Feeds[0].Title)
	}
}

func TestPodcastsDead(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/podcasts/dead" {
			t.Errorf("path = %q, want /podcasts/dead", r.URL.Path)
		}
		w.Write([]byte(`{"status":"true","feeds":[{"id":99,"title":"Dead Feed","dead":1}],"count":1}`))
	}))
	defer srv.Close()

	c := NewClient("key", "secret", WithBaseURL(srv.URL))
	result, err := c.PodcastsDead(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Feeds[0].Dead != 1 {
		t.Errorf("dead = %d, want 1", result.Feeds[0].Dead)
	}
}

func TestPodcastsBatchByGUID(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %q, want POST", r.Method)
		}
		body, _ := io.ReadAll(r.Body)
		if string(body) != `["guid1","guid2"]` {
			t.Errorf("body = %q", string(body))
		}
		w.Write([]byte(`{"status":"true","feeds":[{"id":1,"title":"Feed 1"},{"id":2,"title":"Feed 2"}],"count":2}`))
	}))
	defer srv.Close()

	c := NewClient("key", "secret", WithBaseURL(srv.URL))
	result, err := c.PodcastsBatchByGUID(context.Background(), []string{"guid1", "guid2"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Feeds) != 2 {
		t.Fatalf("got %d feeds, want 2", len(result.Feeds))
	}
}
