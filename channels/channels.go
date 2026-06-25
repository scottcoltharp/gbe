package channels

import (
	"fmt"
)

func Channels() {

	fmt.Println("start")
	messages := make(chan string, 2)

	messages <- "ping"
	messages <- "pong"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
