package api

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"

	"git.mills.io/prologic/bitcask"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Client struct {
	cache    *bitcask.Bitcask
	client   *http.Client
	endpoint string
}

func NewClient(db *bitcask.Bitcask) *Client {
	return &Client{
		cache: db,
		client: &http.Client{
			Timeout: time.Second * 5,
		},
		endpoint: "https://ergast.com/api/f1",
	}
}

func (c *Client) cached(key []byte) (*[]byte, error) {
	if !c.cache.Has(key) {
		logrus.WithFields(logrus.Fields{"key": string(key)}).Info("key not found in cache")

		return nil, bitcask.ErrKeyNotFound
	}
	val, err := c.cache.Get(key)
	if err != nil {
		logrus.WithFields(logrus.Fields{"key": string(key), "error": err}).Error("error getting key value")

		return nil, err
	}
	logrus.WithFields(logrus.Fields{"key": string(key)}).Info("value retrieved from cache")

	return &val, nil
}

func (c *Client) setCacheKey(key, value []byte, ttl *time.Duration) error {
	if nil != ttl {
		if err := c.cache.PutWithTTL(key, value, *ttl); err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err,
				"key":   string(key),
				"val":   string(value),
				"ttl":   ttl.String(),
			}).Error("failed to save key with ttl")

			return err
		}
	} else {
		if err := c.cache.Put(key, value); err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err,
				"key":   string(key),
				"val":   string(value),
			}).Error("failed to save key with ttl")

			return err
		}
	}
	return nil
}

func (c *Client) get(url string, holder any, useCache bool, ttl *time.Duration) error {
	if useCache {
		v, err := c.cached([]byte(url))
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
	if err = c.setCacheKey([]byte(url), tmp, ttl); err != nil {
		logrus.Error("failed to save response to cache")
	}
	err = xml.Unmarshal(tmp, holder)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal xml response")
	}
	return nil
}
