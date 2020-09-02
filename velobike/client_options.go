package velobike

import (
	"net/http"
)

const (
	defaultURL = "http://apivelobike.velobike.ru"
)

// ClientOptions defines possible configuration for Velobike API client.
type ClientOptions struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// Base URL for API requests.
	baseURL string
}

// ClientOption defines a type to provide options.
type ClientOption func(*ClientOptions)

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(client *http.Client) ClientOption {
	return func(opts *ClientOptions) {
		opts.client = client
	}
}

// WithBaseURL sets a custom base URL.
func WithBaseURL(url string) ClientOption {
	return func(opts *ClientOptions) {
		opts.baseURL = url
	}
}
