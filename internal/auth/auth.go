package auth

import (
	"crypto/sha1"
	"fmt"
	"strconv"
	"time"
)

// GenerateHeaders returns the four required authentication headers for the PodcastIndex API.
func GenerateHeaders(apiKey, apiSecret, userAgent string) (map[string]string, error) {
	return GenerateHeadersAt(apiKey, apiSecret, userAgent, time.Now())
}

// GenerateHeadersAt returns auth headers using the given timestamp (useful for testing).
func GenerateHeadersAt(apiKey, apiSecret, userAgent string, t time.Time) (map[string]string, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("API key is required")
	}
	if apiSecret == "" {
		return nil, fmt.Errorf("API secret is required")
	}

	ts := strconv.FormatInt(t.Unix(), 10)
	hash := sha1.Sum([]byte(apiKey + apiSecret + ts))
	authorization := fmt.Sprintf("%x", hash)

	return map[string]string{
		"User-Agent":   userAgent,
		"X-Auth-Key":   apiKey,
		"X-Auth-Date":  ts,
		"Authorization": authorization,
	}, nil
}
