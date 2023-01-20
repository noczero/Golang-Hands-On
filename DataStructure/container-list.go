package main

import (
	"container/list"
	"fmt"
)

/*
Implement double linked list as list package
[nil]->[x]<->[x]<->[nil]
prev, next, pushback

*/

type Dummy struct {
	name string
}

func main() {

	myList := list.New()
	myList.PushBack("One")
	myList.PushBack("Two")
	myList.PushBack("Three") // add data to back

	myList.PushFront(Dummy{name: "test"}) // add data to front

	// how to print using for
	for value := myList.Front(); value != nil; value = value.Next() {
		fmt.Println(value.Value)
	}

}
