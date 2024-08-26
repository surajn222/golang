package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	go simpleGoroutine()
	time.Sleep(1 * time.Second)
}

func simpleGoroutine() {
	fmt.Println("Hello World")
	_, _ = os.Create("newFile.txt")
}
