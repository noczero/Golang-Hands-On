package main

import "fmt"

/*
Closure is an inner function that allows us to interact with outer variable without redefine it.
But be careful to use it, unless understand what to change in outer variable
*/
func main() {
	counter := 0
	name := "Satrya"

	increment := func() {
		fmt.Println("Increment counter in closure")
		counter++
	}

	changeName := func() {
		fmt.Println("Change named")
		name = "Satria"
	}

	scopeClousure := func() {
		fmt.Println("Scope variable cant be access in outer ")
		innerValue := 12345
		fmt.Println(innerValue)
	}
	//innerValue = 3 // wrong

	increment()
	increment()
	fmt.Println(counter)

	changeName()
	fmt.Println(name)

	scopeClousure()
}
