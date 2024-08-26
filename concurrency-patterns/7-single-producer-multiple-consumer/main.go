package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Starting")
	chan1 := make(chan int)

	var wg sync.WaitGroup
	wg.Add(3)

	go producer(&wg, chan1)
	go consumer(&wg, chan1, 1)
	go consumer(&wg, chan1, 2)

	wg.Wait()
	//close(chan1)
}

func producer(wg *sync.WaitGroup, ch chan int) {
	fmt.Println("starting producer")
	defer wg.Done()

	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)
}

func consumer(wg *sync.WaitGroup, ch chan int, consumerid int) {
	defer wg.Done()
	fmt.Println("starting consumer")

	for val := range ch {
		fmt.Println(consumerid, val)
	}

}
