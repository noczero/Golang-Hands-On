package main

import "fmt"

/*
Interface is contract for others method for input or output as parameter.

In GO the implementation of interface is implisit declaration
where the interface with same function name in struct method and return value is same with the contract.
*/

// init interface
type HasName interface {
	// function name | return type
	GetName() string
}

// function that implement interface data type parameter
func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

// declare struct
type Person struct {
	name string
}

// implement interface with same name and same data type return in struct
func (person Person) GetName() string {
	return person.name
}

// others struct
type Pet struct {
	name string
}

// same implement interface
func (pet Pet) GetName() string {
	return pet.name
}

func main() {
	person := Person{name: "Satrya"}
	SayHello(person) // use some function that has interface parameter

	myPet := Pet{name: "Lilo"}
	SayHello(myPet)
}
