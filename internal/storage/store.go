package store

import "sync"

type Store struct {
	mu sync.RWMutex
	data map[string]string
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func (s *Store) Put(key, value string) {
	s.mu.Lock()
	s.data[key] = value
	s.mu.Unlock()
}

func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()
	value, exists := s.data[key]
	s.mu.RUnlock()
	return value, exists	
}

func (s *Store) Delete(key string) {
	s.mu.Lock()
	delete(s.data, key)
	s.mu.Unlock()
}