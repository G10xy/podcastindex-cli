package testutil

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

// NewTestServer creates an httptest.Server with the given handler and registers
// cleanup on t.Cleanup. Returns the server so callers can use srv.URL to
// configure their own client.
func NewTestServer(t *testing.T, handler http.Handler) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)
	return srv
}

// NewTestServerFunc is a convenience wrapper that accepts an http.HandlerFunc.
func NewTestServerFunc(t *testing.T, handler http.HandlerFunc) *httptest.Server {
	t.Helper()
	return NewTestServer(t, handler)
}

// LoadFixture reads a JSON fixture file from the testdata directory.
func LoadFixture(t *testing.T, name string) []byte {
	t.Helper()
	_, thisFile, _, _ := runtime.Caller(0)
	path := filepath.Join(filepath.Dir(thisFile), "testdata", name)
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("loading fixture %q: %v", name, err)
	}
	return data
}
