// Golang program to demonstrate the variables
// storing the hexadecimal values
package main

import "fmt"

func main() {

	// storing the hexadecimal
	// values in variables
	x := 0xFF
	y := 0x9C

	// Displaying the values
	fmt.Printf("Type of variable x is %T\n", x)
	fmt.Printf("Value of x in hexadecimal is %X\n", x)
	fmt.Printf("Value of x in decimal is %v\n", x)

	fmt.Printf("Type of variable y is %T\n", y)
	fmt.Printf("Value of y in hexadecimal is %X\n", y)
	fmt.Printf("Value of y in decimal is %v\n", y)

	// taking a normal variable
	var z int = 5748

	// declaration of pointer
	var p *int

	// initialization of pointer
	p = &z

	// displaying the result
	fmt.Println("Value stored in z = ", z)
	fmt.Println("Address of z = ", &z)
	fmt.Println("Value stored in variable p = ", p)

	// taking a pointer
	var s *int

	// displaying the result
	fmt.Println("s = ", s)
}
