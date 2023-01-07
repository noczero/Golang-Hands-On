package main

import "fmt"

func main() {

	type noKTP string // use alias
	type married bool

	var noKTPalias noKTP = "123123123312"
	var marriedStatus married = false

	fmt.Println(noKTPalias)
	fmt.Println(marriedStatus)
}
