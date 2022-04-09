package locker_guard

import (
	"sync"
)

type MutexGuard[T any] struct {
	mu    sync.Mutex
	value T
}

func NewMutexGuard[T any](value T) MutexGuard[T] {
	return MutexGuard[T]{value: value}
}

func (g *MutexGuard[T]) Do(fn func(v *T)) {
	g.mu.Lock()
	defer g.mu.Unlock()
	fn(&g.value)
}

func (g MutexGuard[T]) Unwrap() T {
	return g.value
}
