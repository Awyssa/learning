package main

import (
	"fmt"
	"unicode/utf8"
);

func main() {
	fmt.Println("Lesson 1: Variables");

	// Numbers
	var myNumber int = 100;
	fmt.Println(myNumber)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	// Strings
	var myString string = "Hello" + " " + "World!"
	fmt.Println(myString)

	// Len will print the number of bytes the string has, not the length.
	// Use unicode/utf8 package for that
	fmt.Println(len("This is x bytes"))
	fmt.Println(utf8.RuneCountInString("This is x characters long"))

	// Runes. I dunno what these are or why its mentioned... Why is is 97?!
	var myRune rune = 'a'
	fmt.Println(myRune)

	// Bool is a boolean, obvi
	var myBoolean bool = false;
	fmt.Println("This boolean is", myBoolean)

	// We can create a variable with no value like...
	var noValueInt int;
	fmt.Println(noValueInt)	
	// for ints it will be 0, string is "", bool is false

	// Go will auto infer the type if not declared, such as
	var oneMoreString = "This is a string"
	fmt.Println(oneMoreString)

	// we can also drop the var and just do
	anotherString := "Another string"
	fmt.Println(anotherString)

	// You can also create multiple values like so...
	var1, var2, var3 := 1, 2, 3;
	fmt.Println(var1, var2, var3)

	// One issue with this is that we do not know what functions return, like so
	value := foo()
	fmt.Println(value)

	// wtf is someValue going to be? We should do this
	var value2 rune = foo();
	fmt.Println(value2)

	// Now we know what type the value will be, yay!

	// we also have const just like JS
	const myConst string = "const value"
	fmt.Println(myConst);
}

func foo() rune {
	var randomRune rune = 'l'
	return randomRune
}
