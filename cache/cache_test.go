package cache

import (
	"io"
	"testing"
	"time"

	"git.mills.io/prologic/bitcask"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestCache_NewCache(t *testing.T) {
	logrus.SetOutput(io.Discard)
	c, err := NewCache("/tmp/NewCache")
	assert.Nil(t, err, "error occurred starting db in /tmp")
	defer func() { _ = c.Close() }()
}

func TestCache_NewCache_Error(t *testing.T) {
	logrus.SetOutput(io.Discard)
	_, err := NewCache("/etc")
	assert.EqualError(t, err, "mkdir /etc/db: permission denied", nil)
}

func TestCache_LookupCached(t *testing.T) {
	logrus.SetOutput(io.Discard)
	key, value := []byte("hello"), []byte("world")
	c, _ := NewCache("/tmp/LookupCached")
	err := c.db.Put(key, value)
	assert.Nil(t, err)
	defer func(c *Cache) {
		_, _ = c.db.Delete(key), c.Close()
	}(c)
	v, err := c.LookupCached(key)
	assert.Nil(t, err, "error getting cache key hello")
	assert.Equal(t, *v, value)
}

func TestCache_LookupCached_Error_NotFound(t *testing.T) {
	logrus.SetOutput(io.Discard)
	key := []byte("hello")
	c, _ := NewCache("/tmp/LookupCached_Error_NotFound")
	defer func() { _ = c.Close() }()
	v, err := c.LookupCached(key)
	assert.Nil(t, v, "got value for non existent key")
	assert.EqualError(t, err, bitcask.ErrKeyNotFound.Error(), "unexpected error")
}

func TestCache_SetCacheKey(t *testing.T) {
	logrus.SetOutput(io.Discard)
	key, value := []byte("hello"), []byte("world")
	c, _ := NewCache("/tmp/TestCache_SetCacheKey")
	defer func(c *Cache) {
		_, _ = c.db.Delete(key), c.Close()
	}(c)
	err := c.SetCacheKey(key, value, nil)
	assert.Nil(t, err)
	assert.True(t, c.db.Has(key))
}

func TestCache_SetCacheKey_TTL(t *testing.T) {
	logrus.SetOutput(io.Discard)
	key, value, ttl := []byte("hello"), []byte("world"), time.Millisecond*2
	c, _ := NewCache("/tmp/TestCache_SetCacheKey_TTL")
	defer func(c *Cache) {
		_, _ = c.db.Delete(key), c.Close()
	}(c)
	err := c.SetCacheKey(key, value, &ttl)
	time.Sleep(time.Millisecond * 3)
	assert.Nil(t, err)
	assert.False(t, c.db.Has(key))
}
