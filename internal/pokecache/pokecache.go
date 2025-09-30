package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]CacheEntry
	mutex   *sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (s *Cache) Add(key string, val []byte) {
	newCache := CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	s.mutex.Lock()
	s.entries[key] = newCache
	s.mutex.Unlock()
}

func (s *Cache) Get(key string) ([]byte, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	entry, ok := s.entries[key]
	return entry.val, ok
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries: make(map[string]CacheEntry),
		mutex:   &sync.Mutex{},
	}

	go cache.reapLoop(interval)

	return cache
}

func (s *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		<-ticker.C

		s.mutex.Lock()
		for key, value := range s.entries {
			if time.Since(value.createdAt) > interval {
				delete(s.entries, key)
			}
		}
		s.mutex.Unlock()
	}
}
