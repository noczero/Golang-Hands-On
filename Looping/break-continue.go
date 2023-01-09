package main

import "fmt"

/*
Break to stop the looping inside
Continue to skip the rest of looping inside to start over again.
*/

func main() {

	// break
	for i := 0; i < 10; i++ {

		if i == 5 {
			break // loop exit
		}

		fmt.Println("Loop", i)
	}

	// continue
	for i := 0; i < 10; i++ {

		if i%2 == 0 {
			fmt.Println("Genap", i)
			continue // loop continue to post statement, and start again
		}
		fmt.Println("Ganjil", i)
	}
}
