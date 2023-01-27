package GoRoutine

import (
	"fmt"
	"testing"
	"time"
)

/*
Rules :
- the function that set to goroutine should not return a value as it can't be captured.
*/

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestRunHelloWorld(t *testing.T) {
	// initate the concurency to the function
	go RunHelloWorld()

	fmt.Println("Called Next or Before depends CPU")

	// wait for 1 seconds to make sure the goroutine finished
	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display : ", number)
}

func TestManyGoRoutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i) // set to async
	}

	// wait 3 seconds
	time.Sleep(3 * time.Second)
}
