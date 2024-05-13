package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	filename := "main.go"

	// Check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("File does not exist")
	} else {
		fmt.Println("File exists")
	}

	cmd := exec.Command("ps", "-e")

	// Execute the command and capture its output
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}

	// Split the output into lines
	lines := strings.Split(string(output), "\n")

	// Print each line (each line corresponds to a process)
	for _, line := range lines {
		fmt.Println(line)
	}
}
