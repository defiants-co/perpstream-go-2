package utils

import (
	"sync"
)

type ConcurrentList[T any] struct {
	mu    sync.RWMutex
	items []T
}

func (cl *ConcurrentList[T]) CopyArray(arr []T) {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	cl.items = append(cl.items[:0], arr...)
}

func (cl *ConcurrentList[T]) GetCopy() []T {
	cl.mu.RLock()
	defer cl.mu.RUnlock()
	copyArr := make([]T, len(cl.items))
	copy(copyArr, cl.items)
	return copyArr
}

func (cl *ConcurrentList[T]) Append(item ...T) {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	cl.items = append(cl.items, item...)
}

func (cl *ConcurrentList[T]) Remove(index int) {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	if index >= 0 && index < len(cl.items) {
		cl.items = append(cl.items[:index], cl.items[index+1:]...)
	}
}

func (cl *ConcurrentList[T]) Get(index int) (T, bool) {
	cl.mu.RLock()
	defer cl.mu.RUnlock()
	if index >= 0 && index < len(cl.items) {
		return cl.items[index], true
	}
	var zero T
	return zero, false
}

func (cl *ConcurrentList[T]) Length() int {
	cl.mu.RLock()
	defer cl.mu.RUnlock()
	return len(cl.items)
}

func NewConcurrentList[T any]() *ConcurrentList[T] {
	return &ConcurrentList[T]{
		items: make([]T, 0),
	}
}
