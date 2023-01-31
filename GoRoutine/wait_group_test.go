package GoRoutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
WaitGroup can be used to wait a goroutine process finished automatically without using sleep.
Using method Add(int) and defer Done() inside the function, and to Wait() on entry function
*/

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done() // called done

	group.Add(1) // add specific group

	fmt.Println("Hello")
	time.Sleep(1 * time.Second) // simulate latency
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{} // create group in pakace sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group) // using goroutine
	}

	group.Wait() // wait all goroutine in group are finished
	fmt.Println("Completed")
}

/*
We can use once.Do(functionname) to set function be called just one time
*/
var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {

		go func() {
			group.Add(1)
			once.Do(OnlyOnce) // this executes one time
			fmt.Println("Executed 100x")
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Conuter : ", counter)
}
