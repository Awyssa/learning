package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Strings with Runes and Bytes (ngl, i'm scared...)")

	// Notice we have a few non ASCII characters
	var myString = "résumé"
	var indexed = myString[0]

	// Notice here we get a number, wtf?
	fmt.Println(indexed)

	// Print the value and type to see what is going on
	fmt.Printf("%v, %T \n", indexed, indexed)

	// We can loop over the string and see further
	for i, v := range myString {
		fmt.Println(i, v)
	}

	// In the print we will see "0 144, 1 233, 3 155, 4 117..." Why did we miss 2?
	// This is because the non ASCII key is 2 bytes, so it takes up 2 spaces
	// NOTE: when using len(), its the length of bytes, not characters

	fmt.Println("--- break ---")

	// This is where runes come in. Rather than couting the bytes, we can cast to runes and see more clearly
	// NOTE: runes are an alias for int32, not int8

	var myString2 = []rune("résumé")

	// We can loop over the string with runes now, and get the more realistic index on 1,2,3,4,5
	for i, v := range myString2 {
		fmt.Println(i, v)
	}

	// If we wanted to split a string, and rebuild it, we'd do something like this (wtf again...)
	var stringSlice = []string{"h", "e", "l", "l", "o"}
	
	var catString = ""

	for i := range stringSlice {
		catString += stringSlice[i]
	}
	fmt.Println(catString)

	// Thankfully, go has the strings package
	var strBuilder strings.Builder

	// Do the catString assignment better
	for i := range stringSlice {
		strBuilder.WriteString(stringSlice[i])
	}

	// Build the string
	var catStr = strBuilder.String()
	fmt.Println(catStr)
}