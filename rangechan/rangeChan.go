package rangechan

import (
	"fmt"
	"time"
)

func RangeChan() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	//queue <- "two"

	for elem := range queue {
		fmt.Println(elem)
	}

	timer1 := time.NewTimer(2 * time.Second)

	timer1.Stop()
}
