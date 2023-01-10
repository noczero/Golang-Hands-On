package main

import "fmt"

/*
Defer is the syntax to set some function to be executed in the end invoke function even there's a error.
Panic is to throw an exception message and stopped the program.
Recover is to catch error message that cause by panic, and define it inside defer function
*/

func logging() {
	fmt.Println("Defer function, logging")
	if errMessage := recover(); errMessage != nil {
		fmt.Println("Error message : ", errMessage)
	}
	fmt.Println("Already executed ")
}

func tryFunction(value int, isPanic bool) {
	defer logging() // define defer in top of function

	// make a panic
	if isPanic {
		panic("Throw an error")
	}

	div := 1 / value
	result := fmt.Sprintf("%.2f", div)
	fmt.Println(result)
}

func main() {
	tryFunction(10, false) // normal case

	tryFunction(10, true) // panic executed case

	tryFunction(0, false) // panic from division by zero

}
