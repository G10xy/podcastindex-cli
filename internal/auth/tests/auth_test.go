package auth

import (
	"crypto/sha1"
	"fmt"
	"testing"
	"time"
)

func TestGenerateHeadersAt(t *testing.T) {
	apiKey := "testkey123"
	apiSecret := "testsecret456"
	userAgent := "podcastindex-cli/test"
	ts := time.Unix(1700000000, 0)

	headers, err := GenerateHeadersAt(apiKey, apiSecret, userAgent, ts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if headers["User-Agent"] != userAgent {
		t.Errorf("User-Agent = %q, want %q", headers["User-Agent"], userAgent)
	}
	if headers["X-Auth-Key"] != apiKey {
		t.Errorf("X-Auth-Key = %q, want %q", headers["X-Auth-Key"], apiKey)
	}
	if headers["X-Auth-Date"] != "1700000000" {
		t.Errorf("X-Auth-Date = %q, want %q", headers["X-Auth-Date"], "1700000000")
	}

	expectedHash := sha1.Sum([]byte(apiKey + apiSecret + "1700000000"))
	expectedAuth := fmt.Sprintf("%x", expectedHash)
	if headers["Authorization"] != expectedAuth {
		t.Errorf("Authorization = %q, want %q", headers["Authorization"], expectedAuth)
	}
}

func TestGenerateHeadersMissingKey(t *testing.T) {
	_, err := GenerateHeaders("", "secret", "ua")
	if err == nil {
		t.Error("expected error for empty API key")
	}
}

func TestGenerateHeadersMissingSecret(t *testing.T) {
	_, err := GenerateHeaders("key", "", "ua")
	if err == nil {
		t.Error("expected error for empty API secret")
	}
}
