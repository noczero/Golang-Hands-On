package main

import (
	"fmt"
	"regexp"
)

/*
Package regexp implements regular expression search
https://pkg.go.dev/regexp

Using RE2 accepted syntax that define here
https://github.com/google/re2/wiki/Syntax
*/
func main() {
	// Compile the expression once, usually at init time.
	// Use raw strings to avoid having to quote the backslashes.
	var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)

	fmt.Println(validID.MatchString("adam[23]"))
	fmt.Println(validID.MatchString("eve[7]"))
	fmt.Println(validID.MatchString("Job[48]")) // false as J is upper case
	fmt.Println(validID.MatchString("snakey"))  // false as not define [ age ]

	re := regexp.MustCompile(`a.`) // regex a and any.

	// find matches of the regex returns a slice of all successive matches.
	fmt.Println(re.FindAllString("test[1] test Abc abc[23]", -1))
	fmt.Println(re.FindAllString("paranormal", 2))
	fmt.Println(re.FindAllString("graal", -1))
	fmt.Println(re.FindAllString("none", -1))
}
