package GoRoutine

import (
	"fmt"
	"strconv"
	"testing"
)

func TestChannelRange(t *testing.T) {
	ch := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- "Looping to " + strconv.Itoa(i)
		}

		// prevent deadlock
		close(ch)
	}()

	// receive data continously until close(channel) is being called
	for data := range ch {
		fmt.Println(data)
	}

	fmt.Println("Finsihed")
}

func TestSelectChannelOnLoop(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- "CH1 : Looping to " + strconv.Itoa(i)
		}

		// prevent deadlock
		close(ch1)
	}()

	go func() {
		for i := 0; i < 10; i++ {
			ch2 <- "CH2 : Looping to " + strconv.Itoa(i)
		}

		// prevent deadlock
		close(ch2)
	}()

	counter := 0
	for {
		// do looping

		select {
		// switch based on channel
		case data := <-ch1:
			fmt.Println(data)
			counter++
		case data := <-ch2:
			fmt.Println(data)
			counter++
		default:
			// default action
			fmt.Println("waiting data")
		}

		// exit stopped when counter is 2, as the data will receive from two channel
		if counter == 2 {
			break
		}
	}

	fmt.Println("Finished")
}
