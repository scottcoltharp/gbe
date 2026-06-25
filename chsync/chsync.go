package chsync

import (
	"fmt"
	"sync"
	"time"
)

func loadTasks(taskCh chan int) {
	for i := range 25 {
		fmt.Println("adding task: ", i)
		time.Sleep(200 * time.Millisecond)
		taskCh <- i
	}
	close(taskCh)
}

func workTasks(id int, taskCh chan int, checkTask chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range taskCh {
		fmt.Printf("Worker #%d working on task: %d\n", id, item)
		time.Sleep(1500 * time.Millisecond)
		checkTask <- item
	}
}

func checkTasks(id int, checkTask chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range checkTask {
		fmt.Printf("Manager #%d checking task: %d\n", id, item)
		time.Sleep(500 * time.Millisecond)
	}
}

func ChSync() {
	taskCh := make(chan int, 30)
	checkTask := make(chan int, 30)
	var workerWg sync.WaitGroup
	var manageWg sync.WaitGroup

	go loadTasks(taskCh)
	for workerID := 1; workerID <= 20; workerID++ {
		workerWg.Add(1)
		go workTasks(workerID, taskCh, checkTask, &workerWg)
	}

	for managerID := 1; managerID <= 10; managerID++ {
		manageWg.Add(1)
		go checkTasks(managerID, checkTask, &manageWg)
	}
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
}
