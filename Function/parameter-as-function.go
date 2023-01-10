package main

import "fmt"

/*
Function can pass in the function parameter
*/

// type declaration for shortening function data type
type Filter func(string) string

func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)

	fmt.Println("Hello", nameFiltered)
}

func sayHelloWithFilterTypeDeclaration(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

// make filter
func spamFilter(name string) string {
	if name == "fuck" {
		return "..."
	} else {
		return name
	}
}

func spamFilterTwo(name string) string {
	if name == "beep" {
		return "---"
	} else {
		return name
	}
}

func main() {
	// we can invoke the function with differenet function as parameter

	sayHelloWithFilter("fuck", spamFilter)
	sayHelloWithFilter("beep", spamFilterTwo)

	sayHelloWithFilterTypeDeclaration("fuck", spamFilter)
	sayHelloWithFilterTypeDeclaration("beep", spamFilterTwo)
}
