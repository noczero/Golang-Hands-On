package main

import "fmt"

func sayHello() {
	fmt.Println("Hello world")
}

func main() {
	sayHello()
	sayHello()
	sayHello()
	sayHello()

	for i := 0; i < 10; i++ {
		fmt.Println(i + 1)
		sayHello()
	}
}
