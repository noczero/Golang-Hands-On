package main

import "fmt"

/*
Map is key-value pair data type, same like dictionary in Python
*/

func main() {
	person := map[string]string{
		"name":    "Satrya",
		"address": "Medan",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	// add values
	person["country"] = "Indonesia"
	fmt.Println(person)

	// usefull function
	fmt.Println(len(person))  // get length
	delete(person, "country") // delete country
	fmt.Println(person)

	// using make to define
	var book map[string]string = make(map[string]string)
	book["title"] = "Book Title"
	book["Author"] = "Book Author"
	book["oops"] = "oops"
	fmt.Println(book)

	delete(book, "oops")
	fmt.Println(book)
}
