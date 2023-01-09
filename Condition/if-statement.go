package main

import "fmt"

func main() {
	// if statement to check condition
	code := 200

	if code == 400 {
		fmt.Println("Bad Request")
	} else if code == 500 {
		fmt.Println("Server Error")
	} else {
		fmt.Println("OK")
	}

	// short if statement, declare variable in if line
	if status := "OK"; status == "OK" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Rejected")
	}
}
