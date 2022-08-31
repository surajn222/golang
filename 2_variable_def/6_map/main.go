// Go program to illustrate how to
// create and initialize maps
package main

import "fmt"

func main() {

	// Method 1
	var map_1 map[int]int
	fmt.Println("map def method 1: ", map_1)

	// Checking if the map is nil or not
	if map_1 == nil {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	// Method 2
	var map_2 = make(map[float64]string)
	fmt.Println("map def method 2: ", map_2)

	// Method 3
	map_3 := map[int]string{
		90: "Dog",
		91: "Cat",
		92: "Cow",
		93: "Bird",
		94: "Rabbit",
	}

	fmt.Println("map def method 3: ", map_3)

	// Iterate
	for k, v := range map_3 {
		fmt.Println(k, v)
	}
}
