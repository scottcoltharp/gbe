package chsync

import (
	"fmt"
	"time"
)

func loadTasks(taskCh chan int) {
	for i := range 25 {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("adding task: ", i)
		taskCh <- i
	}
	close(taskCh)
}

func workTasks(taskCh chan int, done chan bool) {
	for item := range taskCh {
		fmt.Println("working on task: ", item)
		time.Sleep(800 * time.Millisecond)
	}
	done <- true
}

func worker(workdone chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	workdone <- true
}

func manager(workdone chan bool, managedone chan bool) {
	<-workdone
	fmt.Print("Checking employee work...")
	time.Sleep(time.Second)
	fmt.Println("done")
	managedone <- true
}

func ChSync() {
	taskCh := make(chan int, 100)
	done := make(chan bool)

	go loadTasks(taskCh)
	go workTasks(taskCh, done)
	//workdone := make(chan bool)
	//managedone := make(chan bool)
	//go worker(workdone)
	//go manager(workdone, managedone)
	//<-managedone
	<-done
}
