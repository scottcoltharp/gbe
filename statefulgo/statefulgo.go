package statefulgo

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func StatefulGo() {

	var readOps atomic.Uint64
	var writeOps atomic.Uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	for range 100 {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				readOps.Add(1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for range 10 {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(500),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				writeOps.Add(1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := readOps.Load()
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := writeOps.Load()
	fmt.Println("writeOps:", writeOpsFinal)
	//fmt.Println("Current value of key 0:", state[0])
}
