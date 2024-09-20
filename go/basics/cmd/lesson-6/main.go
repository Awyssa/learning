package main

import (
	"fmt"
)

func main() { 
	fmt.Println("Pointers in go")

	// wtf is a pointer?
	//In Go, a pointer is a variable that stores the memory address of another variable. 
	// Instead of holding a value directly, a pointer "points" to the location in memory where the value is stored. 
	// Pointers allow you to reference and manipulate the data at that memory address.

	// Here we are allocation p a memory location of int32 somewhere in memory
	var p *int32 = new(int32)

	// Here we are letting go do the allocation for us
	var i int32

	// Lets look whats going on under the hood (the * is called deferancing the value)
	fmt.Printf("The value p points to is: %v", *p)
	fmt.Printf("\nThe value if i is: %v", i)

	// Here we saying, set the location value of p to this new value
	*p = 10

}