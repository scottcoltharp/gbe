package atomicexample

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func AtomicExample() {

	var ops atomic.Uint64
	var ops2 int = 0

	var wg sync.WaitGroup

	for range 100 {
		wg.Go(func() {
			for range 1000 {
				ops.Add(1)
				ops2++
			}
		})
	}

	wg.Wait()

	fmt.Println("ops:", ops.Load())
	fmt.Println("ops2:", ops2)
}
