// Package cedarclient talks to the cedar-agent sidecar (a small Rust HTTP
// service wrapping AWS's cedar-policy crate) over localhost. This is where
// the "sidecar" architecture decision lives: the Go process never links
// Cedar directly — it pushes the active policy set to the sidecar on every
// change and calls it for each /evaluate request. See sidecar/cedar-agent/README.md.
package cedarclient

import (
	"net/http"
	"time"
)

type Client struct {
	baseURL    string
	authToken  string
	httpClient *http.Client
}

func New(baseURL, authToken string) *Client {
	return &Client{
		baseURL:   baseURL,
		authToken: authToken,
		httpClient: &http.Client{
			// Sidecar is on localhost; a tight timeout keeps a stuck agent
			// from blowing the <5ms evaluation budget's error case out
			// further than necessary — callers should treat a timeout as a
			// fail-closed DENY, not retry inline.
			Timeout: 50 * time.Millisecond,
		},
	}
}
