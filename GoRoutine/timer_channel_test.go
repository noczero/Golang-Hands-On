package GoRoutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
Timer will invoke event in specific time for once and after that it will send data to channel.
It useful for delaying a job or task.
*/

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second) // make timer for 5 seconds

	fmt.Println(time.Now()) // execute first

	time := <-timer.C // store it on time, will invoke 5 seconds later

	fmt.Println(time)
}

func TestAfter(t *testing.T) {
	channel := time.After(5 * time.Second) // make timer for 5 seconds

	fmt.Println(time.Now()) // execute first

	time := <-channel // will invoke 5 seconds later
	fmt.Println(time)
}

func TestAfterFunction(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	// simplfy timer to delay function, will execute 5 seconds later
	time.AfterFunc(5*time.Second, func() {
		fmt.Println(time.Now())
		group.Done()
	})

	fmt.Println(time.Now()) // execute first

	group.Wait()
}
