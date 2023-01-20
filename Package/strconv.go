package main

import (
	"fmt"
	"strconv"
)

/*
Package strconv implements conversions to and from string representations of basic data types.
including
atoi, itoa, format, parse
*/

func main() {

	//  string to integer 64
	val, err := strconv.Atoi("1234")
	if err == nil {
		fmt.Println(val)
	}

	// another string to interger, base : 2 (binary), 10 (decimal), 16(hexadecimal), bitSize : 16,32,64
	val64, err := strconv.ParseInt("1234", 10, 64)
	if err == nil {
		fmt.Println(val64)
	}

	// something to integer
	valStr := strconv.Itoa(12345)
	fmt.Println(valStr)

	valStr = strconv.FormatInt(65423, 10)
	fmt.Println(valStr)

}
