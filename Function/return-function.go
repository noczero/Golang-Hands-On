package main

import "fmt"

// define data type after argument
func calculateResult(name string, value1 float32, value2 float32) string {
	gretting := "Halo " + name
	result := (value1 + value2) * 100
	return gretting + fmt.Sprintf("%.2f", result)
}

// return multiple value
func getFullName() (string, string) {
	return "Satrya", "Budi"
}

// named return values, declare return variable with data type
func getMyName() (firstName string, secondName string, lastName string) {
	firstName = "Satrya"
	secondName = "Budi"
	lastName = "Pratama"

	return
}

func main() {
	// invoke function
	resultStr := calculateResult("Satrya ", 0.51, 0.42)
	fmt.Println(resultStr)

	// invoke function in multiple var
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	// ignore some variable using _
	_, lastNameNew := getFullName()
	fmt.Println(lastNameNew)

	// with named return value
	myFirstName, myMiddleName, myLastName := getMyName()
	fmt.Println(myFirstName, myMiddleName, myLastName)
}
