package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCategoriesList(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/categories/list" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Write([]byte(`{"status":"true","feeds":[{"id":1,"name":"Arts"},{"id":2,"name":"Books"}],"count":2}`))
	}))
	defer srv.Close()

	c := NewClient("key", "secret", WithBaseURL(srv.URL))
	result, err := c.CategoriesList(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Feeds) != 2 {
		t.Fatalf("got %d categories, want 2", len(result.Feeds))
	}
	if result.Feeds[0].Name != "Arts" {
		t.Errorf("name = %q, want Arts", result.Feeds[0].Name)
	}
}
