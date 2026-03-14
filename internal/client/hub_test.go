package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHubPubNotify(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/hub/pubnotify" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.URL.Query().Get("id") != "920666" {
			t.Errorf("id = %q, want 920666", r.URL.Query().Get("id"))
		}
		if r.Header.Get("X-Auth-Key") != "" {
			t.Error("hub endpoints should not send auth headers")
		}
		w.Write([]byte(`{"status":"true","description":"Feed marked for update"}`))
	}))
	defer srv.Close()

	c := NewClient("key", "secret", WithBaseURL(srv.URL))
	result, err := c.HubPubNotify(context.Background(), HubPubNotifyOptions{ID: 920666})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Description != "Feed marked for update" {
		t.Errorf("description = %q", result.Description)
	}
}
