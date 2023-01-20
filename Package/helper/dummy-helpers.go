package helper

import "fmt"

var version = "1.3.0" // can't access outside this file
var VERSION = "1.3.0" // can access outside this file

/*
If function or variable starts with lowercase it can't be access outside this file, otherwise.
its called access modifier.
*/

func SayHello(name string) {
	fmt.Println("Hwllo", name)
}

func sayHello(name string) {
	fmt.Println("Hwllo", name)
}
