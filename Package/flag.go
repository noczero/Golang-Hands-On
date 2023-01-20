package main

/*
Flag is package that implements command-line flag parsing
https://pkg.go.dev/flag
*/
import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "hostname", "Put your database host")
	var user *string = flag.String("user", "admin", "Put your database user")
	var password *string = flag.String("password", "admin", "Put your database password")

	flag.Parse()

	fmt.Println("host", *host)
	fmt.Println("user", *user)
	fmt.Println("password", *password)
}
