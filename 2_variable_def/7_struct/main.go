package main

import "fmt"

// Defining a struct type
type struct_1 struct {
	var_1 string
	var_2 string
	var_3 int
}

func main() {

	// Declaring a variable of a `struct` type
	// All the struct fields are initialized
	// with their zero value
	var var_1 struct_1
	fmt.Println(var_1)

	// Declaring and initializing a
	// struct using a struct literal
	var_struct_1 := struct_1{"Akshay", "Dehradun", 3623572}

	fmt.Println("Address1: ", var_struct_1)

	// Naming fields while
	// initializing a struct
	var_struct_2 := struct_1{var_1: "Anikaa", var_2: "Ballia", var_3: 277001}
	fmt.Println("Address2: ", var_struct_2)

	// Uninitialized fields are set to
	// their corresponding zero-value
	var_struct_3 := struct_1{var_1: "Delhi"}
	fmt.Println("Address3: ", var_struct_3)
}
