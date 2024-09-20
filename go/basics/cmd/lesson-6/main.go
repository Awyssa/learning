package main

import (
	"fmt"
)

/*
	In Go, a pointer is a variable that stores the memory address of another variable.
	Instead of holding a value directly, a pointer "points" to the location in memory where the value is stored.
	Pointers allow you to reference and manipulate the data at that memory address.

	When you declare a variable as a pointer to a type, you use the * symbol to indicate that the variable will store the memory address of a value of the specified type.

	When you have a pointer and you want to access the actual value the pointer is pointing to, you use the * symbol to dereference the pointer.
*/

func main() {
	fmt.Println("Pointers in go")

	// Declare a normal integer variable
	var num int = 42

	// Declare a pointer to an integer, and assign it the address of num
	var ptr *int = &num

	// Print the address stored in the pointer (memory address of num)
	fmt.Println("Pointer address:", ptr)

	// Dereference the pointer to access the value stored at that address
	fmt.Println("Pointer value:", *ptr)

	// Modify the value via the pointer
	*ptr = 100

	// The value of num has changed
	fmt.Println("New value of num:", num)

	// Here we are allocation p a memory location of int32 somewhere in memory
	var p *int32 = new(int32)

	// Here we are letting go do the allocation for us
	var i int32

	// Lets look whats going on under the hood (the * is called dereferencing the value)
	fmt.Printf("The value p points to is: %v", *p)
	fmt.Printf("\nThe value if i is: %v \n", i)

	// Here we saying, set the location value of p to this new value
	*p = 10

	// If we copy a slice without a pointer, we get some strange behavior
	var slice = []int32{1,2,3}
	var sliceCopy = slice

	// Lets try modify the sliceCopy
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)

	// We should get [1,2,4] for both slices (wtf?). This is because slices contain pointers to an underlying array.

	// Here we are passing an array to a function to square each index. The issue is that each time we do this, we create a copy of the array.
	// This is not optimal in terms of speed and performance
	// To fix this, we should pass the reference of thing1 to the square function, and tell the function its a pointer
	// We are printing the memory locations to ensure they are the same
	var thing1 = [5]float64{1,2,3,4,5}
	fmt.Printf("\n The memory location of the thing1 array is :%p", &thing1)
	var result [5]float64 = square(&thing1)
	fmt.Printf("\nThe result if: %v", result)
	fmt.Printf("\nThe value of thing1 is: %v", thing1)
}

func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of thing2 array is %p", thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}
