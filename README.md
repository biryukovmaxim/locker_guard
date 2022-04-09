# MutexGuard
MutexGuard is abstraction of Data with mutex

### Motivation:
1) Library users should not think about mutex unlocking
2) No need to look for places where someone use/don't use the mutex from a struct, everyone have to use now =)

### Example:
```
package main

import (
	"fmt"
	"sync"

	guard "locker_guard"
)

func main() {
	guard := guard.NewMutexGuard(0)
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go guard.Do(func(v *int) {
			defer wg.Done()
			*v++
		})
	}
	wg.Wait()
	fmt.Println(guard.Unwrap())
}
```
