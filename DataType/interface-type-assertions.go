package main

import "fmt"

/**
A type assertion provides access to an interface value's underlying concreate value

It will convert to origin value using .(data-type)
*/

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	// it will be error as int
	//l := i.(int)
	//fmt.Println(l)

	// the recommend way is using switch
	var result interface{} = 13
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is integer")
	default:
		fmt.Println("Unknown value")
	}

}
