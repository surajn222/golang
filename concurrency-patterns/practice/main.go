package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("HW")
	channel := make(chan int)
	go print("Even", channel)
	go print("Odd", channel)
	channel <- 0
	time.Sleep(1 * time.Second)
}

func print(eodd string, channel chan int) {

	for {
		val, ok := <-channel
		if !ok {
			return
		}
		val++
		fmt.Println(eodd, val)
		if val > 9 {
			close(channel)
			return
		}

		channel <- val

	}

}
