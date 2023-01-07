package main

import "fmt"

/*
Constant is used to be on several files, go doesn't complain when unused
constant never update its variable, static
*/

func main() {
	// with data type
	const firstName string = "Satrya Budi"
	const lastName string = "Pratama"

	// without data type
	const value = 100

	// multiple
	const (
		country = "Indonesia"
		height  = 166
		age     = 25
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(country)
}
