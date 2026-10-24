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

