package main

import "fmt"

/*
Slice is part of array and reference from it, the size of slice can change.
In Go slice is mostly use than array.

Slice have pointer, length, and capacity.
To make it define array then slice set pointer.
length is length of the slice.
capacity is the capacity of the slice.

Type of slice :
- array[low:high]
- array[low:]
- array[:high] -> from 0 before high
- array[:] -> from 0 to high
*/

func main() {
	// using ... to automate the length of array
	var months = [...]string{
		"January",
		"Feb",
		"March",
		"April",
		"May",
		"June",
		"July",
		"Agustus",
		"Sept",
		"Oct",
		"Nov",
		"Des",
	}

	// slices from array
	var slices = months[4:7]
	fmt.Println(slices)      // [May June July]
	fmt.Println(len(slices)) // 3
	fmt.Println(cap(slices)) // 8

	// update array values, will impact to the slice, be careful!
	months[5] = "Juni"
	fmt.Println(slices) // [May Juni July]

	// update slice value, will impact to the array, it's each reference
	slices[0] = "Mei"
	fmt.Println(months) // [January Feb March April Mei Juni July Agustus Sept Oct Nov Des]

	// new slices from index 10 (Nov) to the end (Des)
	slices2 := months[10:]
	fmt.Println(slices2) // [Nov Des]

	// using append function to add value, if the capacity is over then it will make new slices and has own array.
	slices3 := append(slices2, "NewMonth")
	fmt.Println(slices3)      // [Nov Des NewMonth]
	fmt.Println(len(slices3)) // 3
	fmt.Println(cap(slices3)) // 4

	// will not impact the slices2 and the months
	slices3[1] = "Desember"
	fmt.Println(slices3) // [Nov Desember NewMonth
	fmt.Println(slices2) // [Nov Des]

	// Be careful, if append is use in the middle of array, it will not make new slices and still reference the array
	slices4 := months[2:4]
	fmt.Println(slices4)               // [March April]
	slices5 := append(slices4, "MEII") // add MEII
	fmt.Println(slices5)               // [March April MEII]
	fmt.Println(months)                // [January Feb March April MEII Juni July Agustus Sept Oct Nov Des]

	// Define proper slice using make that not reference any array.
	newSlices := make([]string, 2, 5) // []type,length,capacity

	// set value
	newSlices[0] = "Satrya"
	newSlices[1] = "Budi"
	fmt.Println(newSlices)
	fmt.Println(len(newSlices))
	fmt.Println(cap(newSlices))

	// Copy slices, make two identic slice
	copySlices := make([]string, len(newSlices), cap(newSlices)) // define the structure
	copy(copySlices, newSlices)                                  // copy origin slice to new
	fmt.Println(copySlices)

	// Slices vs Array, tricky to define
	thisArray := [5]int{1, 2, 3, 4, 5}
	thisSclice := []int{1, 2, 3, 4, 5}
	fmt.Println(thisArray)
	fmt.Println(thisSclice)
}
