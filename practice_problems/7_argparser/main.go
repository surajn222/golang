package main

import (
	//"flag"
	"fmt"
	"github.com/spf13/pflag"
)

func main() {
	//fmt.Print("HW")
	name := pflag.String("name", "suraj", "Name of the person")

	age := pflag.Int("age", 18, "Age of the person")

	var name2 string
	var value bool
	var usage string

	name2 = "bool"
	value = true
	usage = "usage of bool"

	pflag.Bool(name2, value, usage)

	pflag.Parse()

	fmt.Println(*name, *age)
}
