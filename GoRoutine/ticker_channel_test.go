package GoRoutine

import (
	"fmt"
	"testing"
	"time"
)

/*
Ticker is method to execute again and again based on specific time. It useful for routine job execution.
*/

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(3 * time.Second) // 3 seconds events
	defer ticker.Stop()
	done := make(chan bool)

	// make go routine to send stop to channel done for 10 seconds
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()

	// looping to channel
	for {
		select {
		case <-done:
			// when received done then finished
			fmt.Println("Done : ")
			return

		case t := <-ticker.C:
			// handle channel ticker.C
			fmt.Println("Current time", t)
		}
	}

}

func TestTick(t *testing.T) {
	channel := time.Tick(3 * time.Second) // 3 seconds events channel and can't be shutdown

	// receive from channel
	for time := range channel {
		fmt.Println(time)
	}
}
