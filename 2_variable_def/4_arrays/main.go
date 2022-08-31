// In Go, an _array_ is a numbered sequence of elements of a
// specific length.

package main

import "fmt"

func main() {

	// Define an array. default value = 0
	// Method 1
	var a [5]int
	fmt.Println("array def method 1:", a)

	// Method 2
	// Use this syntax to declare and initialize an array
	// in one line.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("array def method 2:", b)

	//set number at index
	a[4] = 100
	fmt.Println("set:", a)

	//get number at index
	fmt.Println("get:", a[4])

	// The builtin `len` returns the length of an array.
	fmt.Println("len:", len(a))

	// Array types are one-dimensional, but you can
	// compose types to build multi-dimensional data
	// structures.
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
