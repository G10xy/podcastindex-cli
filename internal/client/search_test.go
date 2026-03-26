package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/G10xy/podcastindex-cli/internal/testutil"
)

func TestSearchByTerm(t *testing.T) {
	srv := testutil.NewTestServerFunc(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/search/byterm" {
			t.Errorf("path = %q, want /search/byterm", r.URL.Path)
		}
		if q := r.URL.Query().Get("q"); q != "batman" {
			t.Errorf("q = %q, want batman", q)
		}
		if _, ok := r.URL.Query()["clean"]; !ok {
			t.Error("expected clean param to be present")
		}
		if r.Header.Get("X-Auth-Key") == "" {
			t.Error("expected auth headers")
		}
		w.Write(testutil.LoadFixture(t, "search_byterm.json"))
	})

	c := NewClient("key", "secret", WithBaseURL(srv.URL))
	result, err := c.SearchByTerm(context.Background(), SearchByTermOptions{
		Query: "batman",
		Clean: true,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Feeds) != 1 {
		t.Fatalf("got %d feeds, want 1", len(result.Feeds))
	}
	if result.Feeds[0].Title != "Batman University" {
		t.Errorf("title = %q, want Batman University", result.Feeds[0].Title)
	}
}

func TestSearchByTitle(t *testing.T) {
	srv := testutil.NewTestServerFunc(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/search/bytitle" {
			t.Errorf("path = %q, want /search/bytitle", r.URL.Path)
		}
		if q := r.URL.Query().Get("q"); q != "daily" {
			t.Errorf("q = %q, want daily", q)
		}
		w.Write([]byte(`{"status":"true","feeds":[{"id":1,"title":"Everything Everywhere Daily"}],"count":1}`))
	})

	c := NewClient("key", "secret", WithBaseURL(srv.URL))
	result, err := c.SearchByTitle(context.Background(), SearchByTitleOptions{Query: "daily"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Feeds[0].Title != "Everything Everywhere Daily" {
		t.Errorf("title = %q, want Everything Everywhere Daily", result.Feeds[0].Title)
	}
}

func TestSearchByPerson(t *testing.T) {
	srv := testutil.NewTestServerFunc(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/search/byperson" {
			t.Errorf("path = %q, want /search/byperson", r.URL.Path)
		}
		w.Write([]byte(`{"status":"true","items":[{"id":100,"title":"Episode 1","feedTitle":"Some Podcast"}],"count":1}`))
	})

	c := NewClient("key", "secret", WithBaseURL(srv.URL))
	result, err := c.SearchByPerson(context.Background(), SearchByPersonOptions{Query: "adam curry", Max: 5})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Items) != 1 {
		t.Fatalf("got %d items, want 1", len(result.Items))
	}
	if result.Items[0].FeedTitle != "Some Podcast" {
		t.Errorf("feedTitle = %q, want Some Podcast", result.Items[0].FeedTitle)
	}
}

func TestSearchMusicByTerm(t *testing.T) {
	srv := testutil.NewTestServerFunc(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/search/music/byterm" {
			t.Errorf("path = %q, want /search/music/byterm", r.URL.Path)
		}
		if _, ok := r.URL.Query()["aponly"]; !ok {
			t.Error("expected aponly param")
		}
		w.Write([]byte(`{"status":"true","feeds":[{"id":200,"title":"Indie Music"}],"count":1}`))
	})

	c := NewClient("key", "secret", WithBaseURL(srv.URL))
	result, err := c.SearchMusicByTerm(context.Background(), SearchMusicByTermOptions{
		Query:  "indie",
		APOnly: true,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Feeds[0].ID != 200 {
		t.Errorf("id = %d, want 200", result.Feeds[0].ID)
	}
}

func TestSearchByTermMaxParam(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if m := r.URL.Query().Get("max"); m != "10" {
			t.Errorf("max = %q, want 10", m)
		}
		w.Write([]byte(`{"status":"true","feeds":[],"count":0}`))
	}))
	defer srv.Close()

	c := NewClient("key", "secret", WithBaseURL(srv.URL))
	_, err := c.SearchByTerm(context.Background(), SearchByTermOptions{Query: "test", Max: 10})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
