package main

import "fmt"

/*
Using pointer in struct method is recommended to save memory and make it simple to manipulate.
*/

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	man := Man{Name: "Satrya"}
	fmt.Println(man)

	man.Married()

	fmt.Println(man)
}
