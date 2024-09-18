package main

import (
	"fmt"
)

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

// The e part assigns the function to the gasEngine type
func (e gasEngine) milesLeft() uint8 {
	return e.gallons*e.mpg
}

func (e elecEngine) milesLeft() uint8 {
	return e.kwh*e.mpkwh
}

type engine interface { 
	milesLeft() uint8
}

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