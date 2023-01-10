package main

import "fmt"

/*
Anonymous function is function without name
*/

type BlackList func(string) bool

func userRegistration(name string, blackList BlackList) {

	if blackList(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {

	// anonymous function can set to variable
	blackListAnonFunc := func(name string) bool {
		return name == "admin"
	}

	userRegistration("Satrya", blackListAnonFunc)

	// anonymous function in invoke parameter
	userRegistration("admin", func(name string) bool {
		return name == "admin"
	})

}
