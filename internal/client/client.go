package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/G10xy/podcastindex-cli/internal/auth"
)

const (
	defaultBaseURL   = "https://api.podcastindex.org/api/1.0"
	defaultUserAgent = "podcastindex-cli/1.0.0"
)

// Client is an HTTP client for the PodcastIndex API.
type Client struct {
	baseURL    string
	apiKey     string
	apiSecret  string
	userAgent  string
	httpClient *http.Client
}

// Option configures a Client.
type Option func(*Client)

// WithBaseURL sets the API base URL.
func WithBaseURL(url string) Option {
	return func(c *Client) { c.baseURL = url }
}

// WithHTTPClient sets a custom http.Client.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) { c.httpClient = hc }
}

// WithUserAgent sets the User-Agent header value.
func WithUserAgent(ua string) Option {
	return func(c *Client) { c.userAgent = ua }
}

// NewClient creates a new PodcastIndex API client.
func NewClient(apiKey, apiSecret string, opts ...Option) *Client {
	c := &Client{
		baseURL:    defaultBaseURL,
		apiKey:     apiKey,
		apiSecret:  apiSecret,
		userAgent:  defaultUserAgent,
		httpClient: http.DefaultClient,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *Client) doRequest(ctx context.Context, method, path string, params url.Values, body io.Reader, authenticated bool) (*http.Response, error) {
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parsing URL: %w", err)
	}

	if params != nil {
		u.RawQuery = params.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if authenticated {
		headers, err := auth.GenerateHeaders(c.apiKey, c.apiSecret, c.userAgent)
		if err != nil {
			return nil, fmt.Errorf("generating auth headers: %w", err)
		}
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	} else {
		req.Header.Set("User-Agent", c.userAgent)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending request: %w", err)
	}

	return resp, nil
}

func (c *Client) decodeResponse(resp *http.Response, target interface{}) error {
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		switch resp.StatusCode {
		case http.StatusUnauthorized:
			return fmt.Errorf("API returned 401 Unauthorized. Check your API key and secret")
		case http.StatusBadRequest:
			return fmt.Errorf("API returned 400 Bad Request: %s", string(bodyBytes))
		default:
			return fmt.Errorf("API returned %d: %s", resp.StatusCode, string(bodyBytes))
		}
	}

	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		return fmt.Errorf("decoding response: %w", err)
	}

	return nil
}

// addBoolFlag appends a boolean query parameter with an empty value when true.
// The PodcastIndex API uses "parameter shall not have a value" pattern.
func addBoolFlag(params url.Values, key string, val bool) {
	if val {
		params.Set(key, "")
	}
}
