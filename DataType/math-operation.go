package main

import "fmt"

func main() {

	var result = 10 + 10
	fmt.Println(result)

	var a = 10
	var b = 10
	var c = a + b
	var d = a * b

	fmt.Println(c, d)

	// augmented assignment, for shortened to its own variable
	var i = 10
	i += 10
	fmt.Println(i)

	var j = 20
	j -= 1
	fmt.Println(j)

	// unary operator
	// ++ (x+=1) , -- (x-=1)  , - , + , ! (inverse boolean)
	i++
	fmt.Println(i)

	j--
	fmt.Println(j)

}
