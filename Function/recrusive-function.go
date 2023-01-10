package main

import "fmt"

func factorialLoopSolution(value int) int {
	// the idea is times the value to decrement value until base value 1.
	// 5! = 5*4*3*2*1
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

func factorialRecrusiveSolution(value int) int {
	if value == 1 {
		// define base value 1
		return 1
	} else {
		// invoke its function with decrement parameter
		return value * factorialRecrusiveSolution(value-1)
	}
}

func main() {
	fmt.Println(factorialLoopSolution(5))
	fmt.Println(factorialRecrusiveSolution(5))
}
