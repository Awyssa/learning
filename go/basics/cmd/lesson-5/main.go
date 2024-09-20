package main

import (
	"fmt"
)

/*
	In Go, a struct (short for structure) is a composite data type that groups together variables under a single type. These variables, known as fields, can have different types. A struct is particularly useful for representing complex data models that are composed of different fields. Structs can be compared to classes in object-oriented languages, but Go does not support inheritance like some OOP languages do.
*/
type gasEngine struct {
	mpg uint8
	gallons uint8
	owner
}

type elecEngine struct {
	mpkwh uint8
	kwh uint8
}

type owner struct {
	name string
}

// The e part assigns the function to the gasEngine type. This is basically a method on the gasEngine "class"
func (e gasEngine) milesLeft() uint8 {
	return e.gallons*e.mpg
}

// Same as above, but for the elecEngine
func (e elecEngine) milesLeft() uint8 {
	return e.kwh*e.mpkwh
}

/*
	In Go, an interface is a type that defines a set of method signatures but does not provide implementations for those methods. Types that implement all the methods in the interface are said to "satisfy" the interface, even if they do not explicitly declare that they do so. This enables a form of polymorphism where different types can be used interchangeably as long as they implement the same interface.
*/
type engine interface { 
	milesLeft() uint8
}

// We are using the engine interface here, ensuring whatever object is passed to the canMakeIt function has a method of milesLeft and has a uint8 return value
func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

func main() {
	fmt.Println("Lesson 5: Structs and interfaces")

	// Struct: In Go, a struct is a composite data type that groups together variables (also called fields) under a single name.

	// Here we create a var called myEngine that inherits the type of the gasEngine (thing OOP and classes), we can then add the values.
	var myEngine gasEngine = gasEngine{mpg: 25, gallons: 15}
	fmt.Println(myEngine)

	// NOTE: you can also create this with implicit values such as...
	// --> var myEngine gasEngine = gasEngine{25, 15}

	// NOTE: Or add the values like
	// --> myEngine.mpg = 20

	// We can also add another struct by doing this
	var myEngine2 gasEngine = gasEngine{25, 15, owner{"Michael"}}
	fmt.Println(myEngine2.mpg, myEngine2.gallons, myEngine2.name)
	fmt.Printf("Total miles left in tank: %v \n", myEngine.milesLeft())

	// We now have 2 types of engines, and can check if they can make it

	canMakeIt(myEngine, 50)

	var myElecEngine elecEngine = elecEngine{100, 10}
	canMakeIt(myElecEngine, 50)
}