package GoRoutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

/*
In module runtime, it provides a specification of the CPU
like number of CPU, the total of thread, and the total of goroutine that running.
By default, number of GOMAXPROCS is same with number of CPU, so it's no need to manually reconfigure the GOMAXPROCS.
*/

func TestRuntimeSpecs(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU : ", totalCpu)

	runtime.GOMAXPROCS(20) // change to 20 thread
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread", totalThread)

	// test go routine
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine", totalGoroutine) // default is 2, for garabage collector and for running the code.

}
