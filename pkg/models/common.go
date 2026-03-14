package models

// APIResponse wraps the common envelope returned by most endpoints.
type APIResponse struct {
	Status      string `json:"status"`
	Description string `json:"description,omitempty"`
	Count       int    `json:"count,omitempty"`
}

type HubPubNotifyResponse struct {
	APIResponse
}

type AddByFeedURLResponse struct {
	APIResponse
	FeedID int `json:"feedId,omitempty"`
}

type AddByItunesIDResponse struct {
	APIResponse
	FeedID int `json:"feedId,omitempty"`
}
