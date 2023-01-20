package main

import (
	"fmt"
	"math"
)

/*
Package math provides basic constants and mathematical functions.
https://pkg.go.dev/math
Round, Floor, Ceil, Max, Min
*/

func main() {

	fmt.Println(math.Round(1.7))
	fmt.Println(math.Round(1.3))

	fmt.Println(math.Floor(1.7))
	fmt.Println(math.Floor(1.3))

	fmt.Println(math.Max(1.7, 1.2))
	fmt.Println(math.Min(1.7, 1.2))
}
