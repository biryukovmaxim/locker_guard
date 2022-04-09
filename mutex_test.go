package locker_guard

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIncrementor(t *testing.T) {
	guard := NewMutexGuard(0)

	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go guard.Do(func(v *int) {
			defer wg.Done()
			*v++
		})
	}
	wg.Wait()
	require.Equal(t, 10, guard.value)
	require.Equal(t, 10, guard.Unwrap())
}
