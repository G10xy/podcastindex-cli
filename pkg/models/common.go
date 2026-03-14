package models

// APIResponse wraps the common envelope returned by most endpoints.
type APIResponse struct {
	Status      string `json:"status"`
	Description string `json:"description,omitempty"`
	Count       int    `json:"count,omitempty"`
}
