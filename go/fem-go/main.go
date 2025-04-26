package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// Defined struct
	person := Person{Name: "Michael", Age: 30}
	fmt.Printf("This is our person %+v\n", person)

	// Lambda struct
	user := struct {
		name string
		id   int
	}{
		name: "John",
		id:   24,
	}

	fmt.Println(user)

	type Address struct {
		Street string
		City   string
	}

	type Contact struct {
		Name    string
		Address Address
	}

	contact := Contact{
		Name: "Mark",
		Address: Address{
			Street: "123 main street",
			City:   "London",
		},
	}

	fmt.Printf("This contact is %s, and the street is %s, and the city is %s\n", contact.Name, contact.Address.City, contact.Address.Street)

}
