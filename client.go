package sparkpost

import "net/http"

// Client struct to return via a constructor
type Client struct {
	apiURL string
	apiKey string
	http   http.Client
}

// New constructor style pattern for building a sparkpost client
func New(http http.Client, url string, apiKey string) *Client {
	return &Client{apiURL: url, apiKey: apiKey, http: http}
}
