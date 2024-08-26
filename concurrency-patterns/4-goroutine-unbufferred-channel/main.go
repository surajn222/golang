package main

import "fmt"

func main() {
	chan1 := make(chan int)
	go unbufferedChannel(chan1)

	for val := range chan1 {
		fmt.Println(val)
	}
	fmt.Println("done")
}

func unbufferedChannel(ch chan int) {
	fmt.Println("Unbufferred channel")
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
