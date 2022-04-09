package locker_guard

import (
	"sync"
)

type RWMutexGuard[T any] struct {
	mu    sync.RWMutex
	value T
}

func NewRWMutexGuard[T any](value T) RWMutexGuard[T] {
	return RWMutexGuard[T]{value: value}
}

func (g *RWMutexGuard[T]) Read(fn func(v T)) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	fn(g.value)
}

func (g *RWMutexGuard[T]) Write(fn func(v *T)) {
	g.mu.Lock()
	defer g.mu.Unlock()
	fn(&g.value)
}

func (g RWMutexGuard[T]) Unwrap() T {
	return g.value
}
