package main

import "fmt"

/*
Concept :
- 1 variable for 1 data type.
- 1 variable is unique.
*/

func main() {
	// declare
	var name string

	// set
	name = "Satrya Budi Pratama"
	fmt.Println(name)

	// update
	name = "Satrya Budi"
	fmt.Println(name)

	var age int
	age = 25
	fmt.Println(age)

	// instead of var use := for variable declaration
	height := 166
	fmt.Println(height)

	country := "Indonesia"
	fmt.Println(country)

	// bulk declaration / multi declaration
	var (
		firstName = "Satrya Budi"
		lastName  = "Pratama"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
