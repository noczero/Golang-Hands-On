package main

import "fmt"

// define struct
type Customer struct {
	Name, Address string
	Age           int
}

// struct method
func (customer Customer) sayHello() {
	fmt.Println("Hello, my Name is", customer.Name)
}

func (customer Customer) getFive() int {
	return 5
}

func main() {

	// create struct
	var customer Customer
	customer.Name = "Satrya"
	customer.Address = "Medan"
	customer.Age = 25
	fmt.Println(customer)
	fmt.Println("Name " + customer.Name)
	fmt.Println("Address " + customer.Address)
	fmt.Println(customer.Age)

	// sruct literales
	customerStructLiterals := Customer{
		// In GO Land shortcut ctrl + space choose fill all fields
		Name:    "Budi",
		Address: "Bandung",
		Age:     25,
	}
	fmt.Println(customerStructLiterals.Name)
	fmt.Println(customerStructLiterals.Address)
	fmt.Println(customerStructLiterals.Age)

	// direc5
	customerDirect := Customer{"Test", "Ttest", 32}
	fmt.Println(customerDirect)

	// invoke method
	customer.sayHello()
	fmt.Println(customer.getFive())

}
