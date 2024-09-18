package main

import (
	"fmt"
)

func main() {
	fmt.Println("Lesson 3: Arrays, Slices, Maps, and Loops");

	// You can specify the length of the array, and the type of each index. The below will create [0,0,0]
	var intArray [3]int32

	// Print index 0
	fmt.Println(intArray[0])

	// Print index 1 and 2
	fmt.Println(intArray[1:3])

	// Update index 1 to a new value
	intArray[1] = 24
	fmt.Println(intArray[1])

	// int32 is 4 bytes of memory, go allocates 12 bytes. 
	// We can print the memory location using &
	fmt.Println(&intArray[0])
	fmt.Println(&intArray[1])
	fmt.Println(&intArray[2])

	// Another way is like this
	var arrTwo [3]int32 = [3]int32{1,2,3}
	fmt.Println(arrTwo)

	// Or like this
	intArrThree := [3]int32{1,2,3}
	fmt.Println(intArrThree)

	// Or finally like this
	intArrFour := [...]int32{1,2,3}
	fmt.Println(intArrFour)
}