package main

import (
	"fmt"
	"os"
)

// use os package to get the env variable which is already set
func envVariable(key string) string {

	// set env variable using os package
	os.Setenv(key, "dbase")

	// return the env variable using os package
	return os.Getenv(key)
}

func main() {
	// os package
	value := envVariable("name")

	fmt.Printf("os package: %s = %s \n", "name", value)
}
