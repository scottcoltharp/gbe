package chsync

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func loadTasks(taskCh chan<- int) {
	for i := range 200 {
		//time.Sleep(2 * time.Millisecond)
		taskCh <- i
		fmt.Println("loaded task: ", i)
	}
	close(taskCh)
}

func workTasks(id int, taskCh <-chan int, checkTask chan<- int, delayCounter *atomic.Int64) {
	for item := range taskCh {
		time.Sleep(1500 * time.Millisecond)
		select {
		case checkTask <- item:
			fmt.Printf("Worker #%d completed task: %d\n", id, item)
		default:
			checkTask <- item // This blocks safely until a worker frees up a slot
			fmt.Printf("Worker #%d completed task: %d\n", id, item)
			delayCounter.Add(1)
		}
	}
}

func checkTasks(id int, checkTask <-chan int) {
	for item := range checkTask {
		time.Sleep(50 * time.Millisecond)
		fmt.Printf("Manager #%d verified task: %d\n", id, item)
	}
}

func ChSync() {
	// This single line starts the timer and defers the print until main() finishes
	defer func(start time.Time) {
		fmt.Println("Total execution time:", time.Since(start))
	}(time.Now())

	taskCh := make(chan int, 300)
	checkTask := make(chan int, 30)
	var workerWg sync.WaitGroup
	var manageWg sync.WaitGroup

	for managerID := 1; managerID <= 2; managerID++ {
		manageWg.Go(func() {
			checkTasks(managerID, checkTask)
		})
	}

	var delayCounter atomic.Int64
	for workerID := 1; workerID <= 20; workerID++ {
		workerWg.Go(func() {
			workTasks(workerID, taskCh, checkTask, &delayCounter)
		})
	}

	go loadTasks(taskCh)

	go func() {
		workerWg.Wait()
		close(checkTask)
	}()

	//workdone := make(chan bool)
	//managedone := make(chan bool)
	//go worker(workdone)
	//go manager(workdone, managedone)
	//<-managedone
	manageWg.Wait()
	fmt.Println("Total Delays: ", delayCounter.Load())
}
