package main

import "fmt"

/**
Nil is empty value for variable, its recommend to set empty value in interface, function, map, slice, pointer, and channel
if there's no value there for easier checking.
*/

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	myMap := NewMap("Test")
	fmt.Println(myMap)

	// use of nil
	nilMap := NewMap("")
	if nilMap == nil {
		fmt.Println("Empty")
	}
}
