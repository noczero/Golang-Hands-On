package main

import (
	"errors"
	"fmt"
)

/**
Error is an built in interface for creating error message.
It's common pattern to return it on a function, to check failure and success.
*/

// we can make a custom app error struct
type appError struct {
	Error   error
	Message string
	Code    int
}

func checkHttpResponse(message string) (string, appError) {
	if message == "success" {
		return message, appError{
			Error:   nil,
			Message: "success",
			Code:    0,
		}
	} else {
		return message, appError{
			Error:   errors.New("Catch E"),
			Message: "failed",
			Code:    1,
		}
	}

}

func main() {
	res, err := checkHttpResponse("success")
	fmt.Println(res)
	fmt.Println(err)

	res, err = checkHttpResponse("failed")
	fmt.Println(res)
	fmt.Println(err)
}
