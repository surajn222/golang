package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	go goroutine1()
	go goroutine2()
	time.Sleep(1 * time.Second)
}

func goroutine1() {
	fmt.Println("Goroutine 1 running")
	os.Create("NewFile1.txt")
}

func goroutine2() {
	fmt.Println("Goroutine 2 running")
	os.Create("NewFile2.txt")
}
