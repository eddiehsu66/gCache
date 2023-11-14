package core

import (
	"gCache/core/lru"
	"sync"
)

type cache struct {
	lock       sync.Mutex
	lruCache   *lru.Cache
	cacheBytes int64
}

func (c *cache) add(key string, value ByteView) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.lruCache == nil {
		c.lruCache = lru.New(c.cacheBytes, nil)
	}
	c.lruCache.Add(key, value)
}

func (c *cache) get(key string) (v ByteView, ok bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.lruCache == nil {
		return
	}
	if v, ok := c.lruCache.Get(key); ok {
		return v.(ByteView), true
	}
	return
}
