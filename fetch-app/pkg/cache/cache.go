package cache

import "sync"

type MapCache struct {
	mu   sync.RWMutex
	data map[string]interface{}
}

func NewMapCache() *MapCache {
	return &MapCache{
		data: make(map[string]interface{}),
	}
}

func (m *MapCache) Get(key string) interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.data[key]
}

func (m *MapCache) Store(key string, value interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.data[key] = value
}
