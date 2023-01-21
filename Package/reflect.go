package main

import (
	"fmt"
	"reflect"
	"strconv"
)

/*
Package reflect implements run-time reflection, allowing a program to manipulate objects with arbitrary types.
https://pkg.go.dev/reflect
https://go.dev/blog/laws-of-reflection
*/

type Sample struct {
	Name string
	// define structTag for validation, value must string
	Age int `required:"true" max:"100" min:"0"`
}

func IsValid(data interface{}) bool {
	// Simple validation struct field, return true if valid and otherwise

	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := reflect.ValueOf(data).Field(i).Interface()

		// check if required should has value string
		if field.Tag.Get("required") == "true" {
			if value == "" {
				return false
			}
		}

		// check minimal
		conv, err := strconv.Atoi(field.Tag.Get("min"))
		if err == nil {
			if conv > value.(int) {
				return false
			}
		}

		// check maximal
		conv, _ = strconv.Atoi(field.Tag.Get("max"))
		if err == nil {
			if conv < value.(int) {
				return false
			}
		}
	}

	return true
}

func main() {
	sample := Sample{"satrya", 25}
	sampleType := reflect.TypeOf(sample)
	fmt.Println(sampleType)

	// get field object in index 0
	structField := sampleType.Field(0)
	fmt.Println(structField.Name)

	// get field object in index 1
	structField = sampleType.Field(1)
	fmt.Println(structField.Name)

	// get tag value of name required
	structFieldTagRequired := structField.Tag.Get("required")
	fmt.Println(structFieldTagRequired)

	// get min
	structFieldTagMin := structField.Tag.Get("min")
	fmt.Println(structFieldTagMin)

	// get max
	structFieldTagMax := structField.Tag.Get("max")
	fmt.Println(structFieldTagMax)

	// check isValid
	fmt.Println(IsValid(sample))

	// assume failed field
	sample2 := Sample{
		Name: "Satrya",
		Age:  -1,
	}
	fmt.Println(IsValid(sample2))

	sample3 := Sample{
		Name: "",
		Age:  25,
	}
	fmt.Println(IsValid(sample3))

}
