package main

import (
	"fmt"
	"github.com/noczero/Golang-Hands-On/Modules"
)

func main() {
	invokeModule := Modules.SayHello("satrya")
	fmt.Println(invokeModule)
}
