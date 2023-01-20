package main

import (
	"fmt"
	"strings"
)

/*
Methods in strings
https://pkg.go.dev/strings
*/

func main() {

	// Contains
	fmt.Println(strings.Contains("Satrya Budi Pratama", "satrya")) // check string contains, return true or false

	// Split
	fmt.Println(strings.Split("Hello,test,C", ",")) // split string using separator to slice

	// ToLower to change lower case
	fmt.Println(strings.ToLower("LOWER"))

	// ToUpper to change upper case
	fmt.Println(strings.ToUpper("aaaa"))

	// ToTitle to change upper case
	fmt.Println(strings.ToTitle("learn go langs enjoy it"))

	// Trim to clean space in begining and in the end
	fmt.Println(strings.Trim("  satrya  ", " "))

	// Replace words
	fmt.Println(strings.ReplaceAll("gonna replace", "replace", "replaced"))
}
