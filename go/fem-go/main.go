package main

import (
	"fmt"
)

func main() {

	// Control flow
	age := 30

	if age >= 18 {
		fmt.Println("You are an adult!")
	} else if age <= 18 {
		fmt.Println("You are not an adult!")
	} else {
		fmt.Println("Not sure...")
	}

	// Loops
	for i := 0; i < 5; i++ {
		fmt.Println("This is i", i)
	}

	// Array (set length)
	contToTen := [5]int{1, 2, 3, 4, 5}
	fmt.Println(contToTen)

	// Slice (dynamic length)
	fruits := []string{"apple", "orange", "pear"}
	fmt.Println(fruits)

	fruits = append(fruits, "pinapple")
	fmt.Println(fruits)
}
