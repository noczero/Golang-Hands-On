package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

/*
Implement ring
->[nil]->[x]->[x}--

*/

func main() {
	data := ring.New(5)

	// iterate over ring
	for i := 0; i < data.Len(); i++ {
		// to init
		data.Value = "Value " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	// iterate over ring
	data.Do(func(value any) {
		fmt.Println(value)
	})

}
