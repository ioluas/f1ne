// Package cache provides cache wrapper for the application
package cache

import (
	"path/filepath"
	"time"

	"git.mills.io/prologic/bitcask"
	"github.com/sirupsen/logrus"
)

// Cache is wrapper struct around bitcask with helper methods for app/api-clients access patterns
type Cache struct {
	db *bitcask.Bitcask
}

// NewCache returns pointer to new app cache or error if encountered
func NewCache(appPath string) (*Cache, error) {
	db, err := bitcask.Open(filepath.Join(appPath, "db"),
		bitcask.WithMaxKeySize(2_048),
		bitcask.WithMaxValueSize(102_400),
		bitcask.WithSync(true),
	)
	if err != nil {
		logrus.WithError(err).Error("failed to open db")
		return nil, err
	}
	logrus.Trace("db connection open")

	return &Cache{db: db}, nil
}

// LookupCached checks for a value in cache, returns pointer to slice of bytes if exists or error if encountered
func (c *Cache) LookupCached(key []byte) (*[]byte, error) {
	if !c.db.Has(key) {
		logrus.WithFields(logrus.Fields{"key": string(key)}).Info("key not found in cache")

		return nil, bitcask.ErrKeyNotFound
	}
	val, err := c.db.Get(key)
	if err != nil {
		logrus.WithFields(logrus.Fields{"key": string(key), "error": err}).Error("error getting key value")

		return nil, err
	}
	logrus.WithFields(logrus.Fields{"key": string(key)}).Info("value retrieved from cache")

	return &val, nil
}

// SetCacheKey sets a key with value (slice of bytes) for given ttl or without if ttl is nil
func (c *Cache) SetCacheKey(key, value []byte, ttl *time.Duration) error {
	if nil != ttl {
		if err := c.db.PutWithTTL(key, value, *ttl); err != nil {
			logrus.WithError(err).WithFields(logrus.Fields{
				"key": string(key),
				"val": string(value),
				"ttl": ttl.String(),
			}).Error("failed to save key with ttl")

			return err
		}
	} else {
		if err := c.db.Put(key, value); err != nil {
			logrus.WithError(err).WithFields(logrus.Fields{
				"key": string(key),
				"val": string(value),
			}).Error("failed to save key with ttl")

			return err
		}
	}
	return nil
}

// Close closes the wrapped bitcask db object
func (c *Cache) Close() error {
	return c.db.Close()
}
