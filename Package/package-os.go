package main

/*
os is package that interact to host operating system
https://pkg.go.dev/os
*/

import (
	"fmt"
	"os"
)

func main() {
	// retreive path, and CLI arguments
	args := os.Args
	fmt.Println(args)

	// get hostname
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(hostname)
	}

	// get envar from the OS, set it using export in linux or mac
	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")
	fmt.Println(username, password)
}
