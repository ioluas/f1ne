// Package ergast provides http client with external cache for Ergast API
package ergast

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"

	"github.com/ioluas/f1ne/cache"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Client is http client wrapper with a cache component and preset endpoint to Ergast API
type Client struct {
	cache    *cache.Cache
	client   *http.Client
	endpoint string
}

// NewClient returns new http client wrapper with cache db to perform Ergast API calls
func NewClient(db *cache.Cache) *Client {
	return &Client{
		cache:    db,
		client:   &http.Client{Timeout: time.Second * 5},
		endpoint: "https://ergast.com/api/f1",
	}
}

func (c *Client) get(url string, holder any, useCache bool, ttl *time.Duration) error {
	if useCache {
		v, err := c.cache.LookupCached([]byte(url))
		if err == nil {
			err = xml.Unmarshal(*v, holder)
			if err != nil {
				return errors.Wrap(err, "failed to unmarshal cached xml response")
			}
			return nil
		}
	}
	res, err := c.client.Get(url)
	if err != nil {
		return errors.Wrap(err, "could not get result form API endpoint")
	}
	tmp, err := io.ReadAll(res.Body)
	if err != nil {
		return errors.Wrap(err, "could not read response body")
	}
	if err = res.Body.Close(); err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("failed to close response body")
	}
	if err = c.cache.SetCacheKey([]byte(url), tmp, ttl); err != nil {
		logrus.Error("failed to save response to cache")
	}
	err = xml.Unmarshal(tmp, holder)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal xml response")
	}
	return nil
}
