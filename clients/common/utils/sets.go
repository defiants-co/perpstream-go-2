package utils

import (
	"fmt"
	"sync"
)

// ConcurrentSet is a thread-safe set implementation
type ConcurrentSet[T comparable] struct {
	mu    sync.RWMutex
	items map[T]struct{}
}

// NewConcurrentSet creates a new ConcurrentSet
func NewConcurrentSet[T comparable]() *ConcurrentSet[T] {
	return &ConcurrentSet[T]{
		items: make(map[T]struct{}),
	}
}

// Add adds an item to the set
func (s *ConcurrentSet[T]) Add(item T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[item] = struct{}{}
}

// Remove removes an item from the set
func (s *ConcurrentSet[T]) Remove(item T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.items, item)
}

// Contains checks if the set contains an item
func (s *ConcurrentSet[T]) Contains(item T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, exists := s.items[item]
	return exists
}

// String returns a string representation of the set
func (s *ConcurrentSet[T]) String() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var items []T
	for item := range s.items {
		items = append(items, item)
	}
	return fmt.Sprintf("%v", items)
}

// Size returns the number of items in the set
func (s *ConcurrentSet[T]) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.items)
}

// ConcurrentStringSet is a thread-safe set implementation for strings
type ConcurrentStringSet = ConcurrentSet[string]

// NewConcurrentStringSet creates a new ConcurrentStringSet
func NewConcurrentStringSet() *ConcurrentStringSet {
	return NewConcurrentSet[string]()
}
