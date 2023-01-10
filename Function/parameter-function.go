package main

import "fmt"

func sumOperator(value1 int, value2 int) {
	sum := value1 + value2
	fmt.Println("Hasil", sum)
}

func main() {
	sumOperator(1, 2)
}
