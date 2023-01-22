package main

import (
	"fmt"
	"github.com/noczero/Golang-Hands-On/Modules/v2"
)

func main() {
	invokeModule := Modules.SayHello("satrya")
	fmt.Println(invokeModule)
}
