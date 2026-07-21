package main

import (
	"fmt"
	"sync"
)

type Storer interface {
	Push([]byte) (int, error)
	Fetch(int) ([]byte, error)
}

type MemoryStorage struct {
	data [][]byte
	mu   sync.RWMutex
}

type StoreProducerFunc func() Storer

func ReturnNewMemoryStore() *MemoryStorage {
	return &MemoryStorage{}
}

func NewMemoryStorage() Storer {
	return &MemoryStorage{
		data: make([][]byte, 0),
	}
}

func (s *MemoryStorage) Push(data []byte) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = append(s.data, data)
	return len(s.data) - 1, nil
}

func (s *MemoryStorage) Fetch(offset int) ([]byte, error) {
	if offset < 0 {
		return nil, fmt.Errorf("OFFSET CANNOT BE LESS THAN ZERO")
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	if offset > (len(s.data)) {
		return nil, fmt.Errorf("Length Of OFFSET IS GREATER THAN DATA PACKET")
	} else {
		return s.data[offset], nil
	}
}
