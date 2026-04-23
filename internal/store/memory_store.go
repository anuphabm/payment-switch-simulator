package store

import (
	"sync"
	"time"
)

type Record struct {
	Response  []byte
	Timestamp time.Time
}

type MemoryStore struct {
	mu    sync.RWMutex
	store map[string]Record
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		store: make(map[string]Record),
	}
}

// Check if transaction already exists
func (m *MemoryStore) Get(traceID string) (Record, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	rec, ok := m.store[traceID]
	return rec, ok
}

// Save transaction result
func (m *MemoryStore) Set(traceID string, resp []byte) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.store[traceID] = Record{
		Response:  resp,
		Timestamp: time.Now(),
	}
}

func (m *MemoryStore) Cleanup(ttl time.Duration) {
	for {
		time.Sleep(ttl)
		m.mu.Lock()
		for k, v := range m.store {
			if time.Since(v.Timestamp) > ttl {
				delete(m.store, k)
			}
		}
		m.mu.Unlock()
	}
}
