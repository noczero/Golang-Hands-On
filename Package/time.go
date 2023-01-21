package main

import (
	"fmt"
	"time"
)

/*
Package time provides functionality for measuring and displaying time.
https://pkg.go.dev/time

Some constant to define format : RFC822, RFC850, RFC1123, RFC3339
*/

func main() {
	// get date
	fmt.Println(time.Now())
	fmt.Println(time.Date(2023, 01, 21, 7, 30, 10, 1000, time.UTC))

	// convert string to date
	parsed, err := time.Parse(time.RFC3339, "2023-01-21T07:39:05Z")
	if err == nil {
		fmt.Println(parsed)
	} else {
		fmt.Println(err)
	}

}
