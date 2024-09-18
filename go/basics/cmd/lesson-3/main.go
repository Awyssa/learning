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
	intArray[0] = 0
	intArray[1] = 24
	intArray[2] = 42
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

	// Or finally like this. NOTE, the array is still a fixed 3 length
	intArrFour := [...]int32{1,2,3}
	fmt.Println(intArrFour)

	// -- Slice: An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array.
	var intSlice []int32 = []int32{4,5,6}
	fmt.Println(intSlice)

	// Append something to the slice via the append method.
	// NOTE: that when appending, there is an underlying check to see if the initial array has space for the new index. In this case it does not.
	// Therefore, a new array is created in memory.
	fmt.Printf("The length is %v with the capacity of %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("\nThe length is %v with the capacity of %v", len(intSlice), cap(intSlice))

	// NOTE: While the length is now 6, you cannot access these values as you will get a index out of range error
	// This will error out --> fmt.Println(intSlice[5])

	// We can also use the spread operator to add multiple indexes or another slice
	var intSlice2 []int32 = []int32{8,9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	// We can init a splice with the type, current length, and capacity using make
	var intSlice3 []int32 = make([]int32, 3, 6)
	fmt.Println(intSlice3)

	// NOTE: this will just create the slice with the len and cap. You will have to still manually add each index
	intSlice3[0] = 100
	intSlice3[1] = 200
	intSlice3[2] = 300

	fmt.Println(intSlice3)

	// --- Maps: A map maps keys to values (obvi). Basically a JS object, but not the same

	// map = its a map
	// [string] = the key will be a string
	// uint8 = the value will be a uint8

	// NOTE: maps are dynamically sized, so no fixed stuff like arrays
	// NOTE: The zero value of a map is nil, and attempting to write to a nil map will cause a runtime panic.

	var myMap map[string]uint8 = make(map[string]uint8)
	// or: myMap := make(map[string]uint8)
	fmt.Println(myMap)

	// You can also init a map with values, and access them like so
	var myMap2 = map[string]uint8{"Michael": 29, "John": 42}
	fmt.Println(myMap2)
	fmt.Println(myMap2["Json"])

	// If you try to access a key that is not present, you will get back the nil value of the type, so uint8 will be 0
	// fmt.Println(myMap2["score12"]) --> will return 0. So be careful, as a map will always return something even if it doesn't exist

	// Maps can also return an optional second value if the key is not present, this will always be a bool
	var score, ok = myMap2["Json"]
	if ok {
		fmt.Printf("The score is %v", score)
	} else {
		fmt.Println("Invalid score")
	}

	// -- Loops: Go has only one looping construct, the for loop.
	// NOTE: Since maps do not have an order, this will not be in a random order
	for name, age:= range myMap2 {
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}

	// Like maps, we can loop over arrays, where i is the index, and v is the value
	for i, v := range intArray {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}

	// To do a while loop, you can do this
	var i int = 10
	for i >= 0 {
		if i > 0 {
			fmt.Printf("Countdown! %v \n", i)
			i = i - 1
		} else {
			fmt.Println("Happy new year!")
			i -- // can also use --, ++
		}
	}

	// Classic for loop syntax does apply, such as...
	for i := 0; i <= 10; i++ {
		fmt.Printf("Counting to 10, current count is = %v \n", i)
	}
}