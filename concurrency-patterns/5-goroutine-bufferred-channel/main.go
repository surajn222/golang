package main

// Pending
import (
	"fmt"
)

func main() {
	chan1 := make(chan int, 5)

	chan1 <- 1
	chan1 <- 2

	go goroutineChannel(chan1)
	chan1 <- 3
	chan1 <- 4
	chan1 <- 5

	for i := 0; i < 10; i++ {
		chan1 <- i
	}
	// time.Sleep(1 * time.Second)
}

func goroutineChannel(ch chan int) {
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
