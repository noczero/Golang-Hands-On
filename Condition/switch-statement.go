package main

import "fmt"

func main() {

	command := 0

	// switch statement
	switch command {
	case 0:
		fmt.Println("Execute : ", command)
		fmt.Println("Finished")

	case 1:
		fmt.Println("Execute : ", command)
		fmt.Println("Finished")

	default:
		fmt.Println("Unknown command")
	}

	// using short statement in switch
	switch sum := 5 + 4; sum {
	case 9:
		fmt.Println("Easyy")
	}

	// switch without expression, just like if statement but recommend to use on simple condition.
	sum := 7
	switch {
	case sum > 10:
		fmt.Println("High")

	case sum > 5:
		fmt.Println("Normal")

	default:
		fmt.Println("Low")
	}

}
