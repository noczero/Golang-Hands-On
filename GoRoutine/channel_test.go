package GoRoutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {

	// make channel with data type string
	channel := make(chan string)
	defer close(channel) // close channel to avoid memory leaks using defer that executed in the end of the code even error

	// anonymous function without parameter
	go func() {
		time.Sleep(2 * time.Second)
		// send data to channel
		channel <- "Test sent to channel"

	}()

	// receive data from channel
	data := <-channel
	fmt.Println(data)

	time.Sleep(1 * time.Second)
}

// channel as parameter, no need pointer or reference
func GiveMeResponse(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Test channel as parameter" // send daata to ch
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	// receive from channel
	data := <-channel
	fmt.Println(data)

	time.Sleep(2 * time.Second)

}

// to define that function only send to the channel use parameter chan<-
func OnlyIn(ch chan<- string) {
	// send data
	time.Sleep(2 * time.Second)
	ch <- "Try to send using function parameter notation"
}

// to define that function only receive data froo the channel use parameter <-chan
func OnlyOut(ch <-chan string) {
	// receive data
	data := <-ch
	fmt.Println(data)
}

func TestOnlyInOut(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(2 * time.Second)
}

func TestMakeBufferedChannel(t *testing.T) {
	// make buffer channel, set second parameter to total of buffer
	bufferedCh := make(chan string, 3)
	fmt.Println(cap(bufferedCh))
	defer close(bufferedCh)

	go func() {
		// set data based on total buffered
		bufferedCh <- "Send 1"
		bufferedCh <- "Send 2"
		bufferedCh <- "Send 3"
	}()

	go func() {
		// receive data based on set
		fmt.Println(len(bufferedCh))
		fmt.Println(<-bufferedCh)
		fmt.Println(<-bufferedCh)
		fmt.Println(<-bufferedCh)
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Finished")
}
