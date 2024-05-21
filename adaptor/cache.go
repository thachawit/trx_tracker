package adaptor

import (
	"sync"
)

type CacheAdaptor struct {
	Mu    sync.Mutex
	Store map[string]map[string]string
}

func NewCacheService(cache *CacheAdaptor) Cache {
	return &CacheAdaptor{
		Store: cache.Store,
	}
}
func NewCacheAdaptor() *CacheAdaptor {
	cache := &CacheAdaptor{
		Store: make(map[string]map[string]string),
	}
	// Pre-create cache entries if needed
	// Example:
	// cache.Store["eth"] = "hash_transaction"
	return cache
}

func (c *CacheAdaptor) Get(walletAddress string) (map[string]string, bool) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	val, ok := c.Store[walletAddress]
	return val, ok
}

func (c *CacheAdaptor) CompareCacheValue(val map[string]string, symbol string, hash string) bool {
	value, found := val[symbol]
	if !found {
		return true
	}
	if value == hash {
		return false
	}
	return true

}

func (c *CacheAdaptor) UpdateTransaction(val map[string]string, walletAddress string, symbol string, hash string) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	// Retrieve existing transactions for the wallet address

	if val == nil {
		val = make(map[string]string)
	}
	// Update or add the new transaction
	val[symbol] = hash
	c.Store[walletAddress] = val

}
