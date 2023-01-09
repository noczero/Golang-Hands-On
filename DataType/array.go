package main

import "fmt"

/*
Array declaration use the static specific length,
if data exceed to the limit array limit then redefine again with more length
*/

func main() {
	var names [3]string // array string declaration
	names[0] = "Satrya"
	names[1] = "Budi"
	names[2] = "Pratama"

	fmt.Println(names)
	fmt.Println(names[2])

	var values [3]int // array integer declaration
	values[0] = 1
	values[1] = 2
	values[2] = 3
	fmt.Println(values)
	fmt.Println(values[2])

	// array direct assignment
	var arrayValues = [3]float32{
		1.3,
		2.4,
		3.5,
	}

	fmt.Println(arrayValues)

	// get length of array, included the empty values
	fmt.Println(len(arrayValues))

	var emptyArrayList [10]int
	fmt.Println(len(emptyArrayList))
}
