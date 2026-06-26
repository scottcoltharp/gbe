package rangechan

import "fmt"

func RangeChan() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	//queue <- "two"

	for elem := range queue {
		fmt.Println(elem)
	}
}
