package application

import (
	"crypto/sha256"
	"encoding/hex"
	"sync"
	"time"
)

type CacheEntry struct {
	Key       string
	Data      []byte
	CreatedAt time.Time
	Hits      uint64
}
type Cache struct {
	mu    sync.RWMutex
	items map[string]CacheEntry
	limit int
}

func NewCache(limit int) *Cache { return &Cache{items: map[string]CacheEntry{}, limit: limit} }
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	v, ok := c.items[key]
	if !ok {
		return nil, false
	}
	v.Hits++
	c.items[key] = v
	// Return a copy so callers cannot mutate the cache's internal buffer, and so
	// two concurrent Get callers never share a slice. Returning v.Data directly
	// let one caller's edit leak into every other reader and into the stored
	// entry itself.
	out := make([]byte, len(v.Data))
	copy(out, v.Data)
	return out, true
}
func (c *Cache) Put(key string, data []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.limit > 0 && len(c.items) >= c.limit {
		for k := range c.items {
			delete(c.items, k)
			break
		}
	}
	// Take ownership of the bytes: a caller that reuses or mutates its source
	// slice after Put must not be able to change what the cache stored.
	owned := make([]byte, len(data))
	copy(owned, data)
	c.items[key] = CacheEntry{Key: key, Data: owned, CreatedAt: time.Now().UTC()}
}
func CacheKey(parts ...string) string {
	h := sha256.New()
	for _, p := range parts {
		h.Write([]byte{0})
		h.Write([]byte(p))
	}
	return hex.EncodeToString(h.Sum(nil))
}
