package main

import "fmt"

/*
Pointer is used to be memory save to huge variable as its reference to memory.
2 operator that use is * and & to define the pointer.
the * use to manipulate the variable, while & is to define a pointer in varialbe.
*/

type Address struct {
	City, Province, Country string
}

func main() {

	// define a pointer
	var address1 *Address = &Address{
		City:     "Medan",
		Province: "North Sumatera",
		Country:  "Indonesia",
	}
	fmt.Println(address1) // &{Medan North Sumatera Indonesia}

	// define a empty pointer
	var address2 *Address = new(Address)
	fmt.Println(address2) // &{  }

	// point the address 2 to address 1
	address2 = address1
	fmt.Println(address2) // &{Medan North Sumatera Indonesia}

	// modify the address2 value directly will change the address 1 too, as its same reference address
	address2.City = "Bandung"
	address2.Province = "West Java"
	fmt.Println(address2) // &{Bandung West Java Indonesia}
	fmt.Println(address1) // &{Bandung West Java Indonesia}

	// same with this, will update the address 1 too
	*address2 = Address{
		City:     "Medan",
		Province: "North Sumatera",
		Country:  "Indonesia",
	}
	fmt.Println(address2) // &{Medan North Sumatera Indonesia}
	fmt.Println(address1) // &{Medan North Sumatera Indonesia}

	// use new declaration of pointer so address 1 and 2 will has diff reference
	address2 = &Address{
		City:     "Bandung",
		Province: "West Java",
		Country:  "Indonesia",
	}
	fmt.Println(address2) // &{Bandung West Java Indonesia}
	fmt.Println(address1) // &{Medan North Sumatera Indonesia}

	// let say has address 3

}
