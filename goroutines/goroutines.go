package goroutines

import (
	"fmt"
	"sync"
	"time"
)

func f(from string) {
	for i := range 3 {
		time.Sleep(1 * time.Second)
		fmt.Println(from, ":", i)
	}
}

func GoRoutines() {
	var wg sync.WaitGroup

	f("direct")

	wg.Go(func() {
		f("goroutine")
	})

	wg.Go(func() {
		time.Sleep(500 * time.Millisecond)
		f("going")
	})

	wg.Wait()
	fmt.Println("done")
}
