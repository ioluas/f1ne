// Package wikimedia provides http client wrapper for Wikimedia API usage
package wikimedia

import (
	"net/http"
	"time"

	"github.com/ioluas/f1ne/cache"
)

// Client is http client wrapper with a cache component and preset endpoint to Wikimedia API
type Client struct {
	cache    *cache.Cache
	client   *http.Client
	endpoint string
}

// NewClient returns new http client wrapper with cache db to perform Wikimedia API calls
func NewClient(db *cache.Cache) *Client {
	return &Client{
		cache:    db,
		client:   &http.Client{Timeout: time.Second * 5},
		endpoint: "https://en.wikipedia.org/w/api.php",
	}
}
