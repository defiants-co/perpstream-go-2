package utils

import (
	"sync"
)

// ConcurrentMap is a thread-safe map that stores any type.
type ConcurrentMap[T any] struct {
	internalMap map[string]T
	mu          sync.Mutex
}

// NewConcurrentMap creates a new instance of ConcurrentMap.
func NewConcurrentMap[T any]() *ConcurrentMap[T] {
	return &ConcurrentMap[T]{
		internalMap: make(map[string]T),
	}
}

// Values returns a copy of the underlying map.
func (cm *ConcurrentMap[T]) Values() map[string]T {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	copyMap := make(map[string]T, len(cm.internalMap))
	for key, value := range cm.internalMap {
		copyMap[key] = value
	}
	return copyMap
}

// Get retrieves the value associated with the given key.
// It returns the value and a boolean indicating whether the key exists.
func (cm *ConcurrentMap[T]) Get(key string) (T, bool) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	value, exists := cm.internalMap[key]
	return value, exists
}

// Set associates the given value with the given key.
func (cm *ConcurrentMap[T]) Set(key string, value T) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.internalMap[key] = value
}

// Contains checks if the given key exists in the map.
func (cm *ConcurrentMap[T]) Contains(key string) bool {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	_, exists := cm.internalMap[key]
	return exists
}

// SetAll replaces the internal map with the given map.
func (cm *ConcurrentMap[T]) SetAll(newMap map[string]T) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.internalMap = newMap
}

// Delete removes the value associated with the given key.
func (cm *ConcurrentMap[T]) Delete(key string) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	delete(cm.internalMap, key)
}

// Iterate calls the given function for each key-value pair in the map.
func (cm *ConcurrentMap[T]) Iterate(fn func(key string, value T)) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	for key, value := range cm.internalMap {
		fn(key, value)
	}
}

func GetKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

func AppendMaps[K comparable, V any](map1, map2 map[K]V) map[K]V {
	// Create a new map to avoid modifying the original map1
	mergedMap := make(map[K]V)

	// Copy all entries from map1 to the mergedMap
	for key, value := range map1 {
		mergedMap[key] = value
	}

	// Copy all entries from map2 to the mergedMap, overwriting any duplicates
	for key, value := range map2 {
		mergedMap[key] = value
	}

	return mergedMap
}

func ReverseMap[K, V comparable](originalMap map[K]V) map[V]K {
	reversedMap := make(map[V]K)
	for key, value := range originalMap {
		reversedMap[value] = key
	}
	return reversedMap
}
