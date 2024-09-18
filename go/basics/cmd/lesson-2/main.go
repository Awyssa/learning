package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Lesson 2: Functions and Control Flow");

	// Passing an argument to a function. We declare the type with the parameter.
	printValue := "Hello!"
	printMe(printValue)

	// Creating a new variable with the returned value of a function
	var numerator int = 32
	var denominator int = 3
	var result, remainder, err = intDivision(numerator, denominator);

	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v", result)
	} else {
		fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}


// We can declare two returns by using parens and adding the return types
func intDivision(numerator int, denominator int) (int, int, error) {
	var err error;

	if denominator == 0 {
		err = errors.New("cannot divide by zero")

		// We still have to return the two ints, but add the error
		return 0, 0, err
	}

	var result int = numerator / denominator;
	var remainder int = numerator % denominator

	return result, remainder, err
}