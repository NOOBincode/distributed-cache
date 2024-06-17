package MiniCache

import (
	"Distributed_cache/MiniCache/lru"
	"sync"
)

type MiniCache struct {
	mu         sync.Mutex
	lru        *lru.Cache
	cacheBytes int64
}

//使用锁再次封装LRU,使其能够满足并发环境

func (cache *MiniCache) add(key string, value ByteView) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	if cache.lru == nil {
		cache.lru = lru.New(cache.cacheBytes, nil)
	}
	cache.lru.Add(key, value)
}

func (cache *MiniCache) get(key string) (value ByteView, ok bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	if cache.lru == nil {
		return
	}
	if value, ok := cache.lru.Get(key); ok {
		return value.(ByteView), true
	}
	return
}
