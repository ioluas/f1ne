package api

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"time"

	"git.mills.io/prologic/bitcask"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Client struct {
	cache  *bitcask.Bitcask
	client *http.Client
}

func NewClient(db *bitcask.Bitcask) *Client {
	return &Client{
		cache: db,
		client: &http.Client{
			Timeout: time.Second * 5,
		},
	}
}

func (c *Client) getKey(key []byte) (*[]byte, error) {
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

func (c *Client) setKey(key, value []byte, ttl *time.Duration) error {
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

func (c *Client) get(url string, holder any) error {
	v, err := c.getKey([]byte(url))
	if err == nil {
		err = xml.Unmarshal(*v, holder)
		if err != nil {
			return errors.Wrap(err, "failed to unmarshal cached xml response")
		}
	} else {
		res, err := c.client.Get(url)
		if err != nil {
			return errors.Wrap(err, "could not get result form API endpoint")
		}
		tmp, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return errors.Wrap(err, "could not read response body")
		}
		if err = res.Body.Close(); err != nil {
			logrus.WithFields(logrus.Fields{"error": err}).Error("failed to close response body")
		}
		if err = c.setKey([]byte(url), tmp, nil); err != nil {
			logrus.Error("failed to save response to cache")
		}
		err = xml.Unmarshal(tmp, holder)
		if err != nil {
			return errors.Wrap(err, "failed to unmarshal xml response")
		}
	}
	return nil

	// res, err := c.client.Get(url)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "could not get result form API endpoint")
	// }
	// tmp, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "could not read response body")
	// }
	// if err = res.Body.Close(); err != nil {
	// 	logrus.WithFields(logrus.Fields{"error": err}).Error("failed to close response body")
	// }
	// if err = c.setKey([]byte(url), tmp, nil); err != nil {
	// 	logrus.Error("failed to save response to cache")
	// }
	//
	// return &tmp, nil
}
