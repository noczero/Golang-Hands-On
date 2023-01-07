package main

import "fmt"

func main() {
	var intVal32 int32 = 7722

	// casting
	var intVal16 int16 = int16(intVal32)

	// up casting
	var intVal64 = int64(intVal16)

	fmt.Println(intVal16)
	fmt.Println(intVal64)

	// convert byte to string
	var name = "satrya"
	var firstCharByte byte = name[0]
	var firstCharStr string = string(firstCharByte)
	fmt.Println(firstCharStr)
}
