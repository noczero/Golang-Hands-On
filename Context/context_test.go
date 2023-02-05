package Context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

/*
Many functions in Go use the context package to gather additional information about the environment they’re being executed in,
and will typically provide that context to the functions they also call.

Context implements parents and child mechanism,
has 3 main method, WithValue, WithCancel, and WithDeadline.

Cancel signal will invokde in channel ctx.Done()
*/

func TestContext(t *testing.T) {

	// create empty context Background
	background := context.Background()
	fmt.Println(background) // will print context.Background

	// create empty context TODO
	todo := context.TODO()
	fmt.Println(todo) // will print context.TODO

}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	// make child from context A
	contextB := context.WithValue(contextA, "key", "value")
	contextC := context.WithValue(contextA, "c", "C")

	// make child from context B
	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	// make child from context C
	contextF := context.WithValue(contextC, "f", "F")

	// make child from context F
	contextG := context.WithValue(contextF, "g", "G")

	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	// Get value from context
	fmt.Println(contextG.Value("f")) // F
	fmt.Println(contextG.Value("g")) // G
	fmt.Println(contextE.Value("c")) // <nil>

}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination) // close channel

		counter := 1
		// iteration send counter forever, it will goroutine leaks, like zombie that not be used.
		//for {
		//	destination <- counter
		//	counter++
		//}

		// using context done to stop iteration
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++

				// simulate latency
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {

	fmt.Println("total goroutine", runtime.NumGoroutine()) // 2

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)                      // passing the context
	fmt.Println("total goroutine", runtime.NumGoroutine()) // 3

	for n := range destination {
		fmt.Println("Counter", n)

		if n == 10 {
			break
		}
	}
	cancel() // to cancel in context
	time.Sleep(2 * time.Second)

	fmt.Println("total goroutine", runtime.NumGoroutine()) // should 2, not 3.
}

func TestContextWithTimeout(t *testing.T) {

	fmt.Println("total goroutine", runtime.NumGoroutine()) // 2

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second) // make context with timeout 5 seconds, if more 5 seconds then automatically send Done() signal.

	defer cancel() // for less than 5 seconds

	destination := CreateCounter(ctx)                      // passing the context
	fmt.Println("total goroutine", runtime.NumGoroutine()) // 3

	for n := range destination {
		fmt.Println("Counter", n)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("total goroutine", runtime.NumGoroutine()) // should 2, not 3.
}

func TestContextWithDeadline(t *testing.T) {

	fmt.Println("total goroutine", runtime.NumGoroutine()) // 2

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second)) // make context with deadline 5 seconds from now, if more 5 seconds then automatically send Done() signal.

	defer cancel() // for less than 5 seconds

	destination := CreateCounter(ctx)                      // passing the context
	fmt.Println("total goroutine", runtime.NumGoroutine()) // 3

	for n := range destination {
		fmt.Println("Counter", n)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("total goroutine", runtime.NumGoroutine()) // should 2, not 3.
}
