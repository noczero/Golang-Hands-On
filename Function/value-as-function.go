package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	sayGoodBye := getGoodBye // function as variable, can pass to the function, first class citizen (tidak dianak tirikan).

	result := sayGoodBye("Satrya")
	fmt.Println(result)
	fmt.Println(getGoodBye("Satrya "))
}
