package main

import (
	"fmt"
	"sort"
)

/*
Sort package implements sorting to slice using interface
https://pkg.go.dev/sort
*/

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// Implements sort.Interface for [] Person based on the Age field
type ByAge []Person

// 3 methods, Len,Swap,Less
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	fmt.Println(people)

	// 1: invoke sort.Sort
	sort.Sort(ByAge(people))
	fmt.Println(people)

	// 2. Using closure
	sort.Slice(people, func(i, j int) bool {
		// Descending sort
		return people[i].Age > people[j].Age
	})
	fmt.Println(people)

}
