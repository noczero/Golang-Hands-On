package main

import "fmt"

/**
Support any data type with this type of interface for function parameter or return.
*/

func emptyInterface(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return '2'
	} else {
		return false
	}
}

func emptyInterface2(any interface{}) interface{} {
	if any == 1 {
		return 1
	} else if any == 2 {
		return '2'
	} else {
		return false
	}
}

func main() {
	var data interface{} = emptyInterface(1)
	fmt.Println(data)

	nilInterface := emptyInterface(3)
	fmt.Println(nilInterface)

	anyy := emptyInterface2(3)
	fmt.Println(anyy)
}
