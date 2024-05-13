package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("HW")

	file, _ := os.Create("testlog.log")
	file.WriteString("Hello World")
	file.Close()

	// Append to the same file
	fileappend, _ := os.OpenFile("testlog.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	fileappend.WriteString("\nAppended String")
	fileappend.Close()

	// Open the above file
	contents, _ := ioutil.ReadFile("testlog.log")
	fmt.Println(string(contents))

}
