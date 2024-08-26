package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(2)
	go goRoutine1(wg)
	go goRoutine2(wg)

	wg.Wait()
}

func goRoutine1(wg sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("goRoutine1")
}

func goRoutine2(wg sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("goRoutine2")
}
