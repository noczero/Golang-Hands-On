package main

import "fmt"

/*
In Go loop just using for, no while, no do while.
*/

func main() {
	counter := 1

	// simple for
	for counter <= 10 {
		fmt.Println("Loop", counter)
		counter++
	}

	// for with statement consist of init statement (execute once), condition, and post statement (execute in every end loop)
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Loop", counter)
	}

	// iterate over slices or arrays
	slices := []string{"Satu", "Dua", "Tiga", "Empat", "Lima"}

	// classic iteration, with index
	for i := 0; i < len(slices); i++ {
		fmt.Println(slices[i])
	}

	// using for range
	for index, value := range slices {
		fmt.Println("index", index, "=", value)
	}

	// iterate over maps
	person := make(map[string]string)
	person["name"] = "Satrya"
	person["age"] = "25"

	for key, value := range person {
		fmt.Println(key, ":", value)
	}
}
