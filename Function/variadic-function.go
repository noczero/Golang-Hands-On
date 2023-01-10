package main

import "fmt"

/*
Variadic function allows slice parameter to pass inside arguments
*/
func sumAll(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}

	return total
}

func main() {
	total := sumAll(10, 20, 30, 40, 50)
	fmt.Println(total)

	// using slice
	slice := []int{10, 30, 0, 40, 50}
	total = sumAll(slice...)
	fmt.Println(total)
}
