package main

import (
	"fmt"
)

func main() {
	fmt.Println("Lesson 9: Generics")

	/*
		Generics in Go, introduced in Go 1.18 (released in March 2022), are a feature that allows you to write code that can work with different types while maintaining type safety. They provide a way to create functions, types, and methods that can operate on various data types without sacrificing compile-time type checking.
	*/

	var intSlice = []int{1, 2, 3}
	// Call sumSlice with int type and print the result
	fmt.Println(sumSlice[int](intSlice))

	var float32Slice = []float32{1, 2, 3}
	// Call sumSlice with float32 type and print the result
	fmt.Println(sumSlice[float32](float32Slice))
}

// Define a generic function sumSlice
// T is a type parameter constrained to int, float32, or float64
// The function takes a slice of T and returns a T
func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}

	return sum
}
