package GoRoutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
Race condition occurs when 1 variable that manipulate by multiple goroutine,
to prevent this go provides lock and unlock mechanism using sync.Mutex (Mutual Exclusion) or sync.RWMutex between variable
or using sync.Puul for simply operational using get and set to the pool.
or using sync.Map to store the variable.
or using atomic package to manipulate single primitive data type.

*/

func TestRaceConditionDemo(t *testing.T) {
	// this variable will share to multiple goroutine, it makes the value will be skipped when the time is same for some goroutine.
	x := 0

	for i := 1; i <= 1000; i++ {

		// goroutine
		go func() {
			for j := 1; j <= 100; j++ {
				x = x + 1 // variable race condition
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter : ", x) // prints less than 100000
}

func TestMutex(t *testing.T) {
	// Mutual Exclusion that set lock and unlock over variable will handle race condition

	x := 0
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {

		// goroutine
		go func() {
			for j := 1; j <= 100; j++ {

				mutex.Lock() // start locking

				x = x + 1 // variable race condition

				mutex.Unlock() // unlock
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter : ", x) // prints 100000
}

/*
## RWMutex
	to handle Read and Write in variable simultaneously we need two mutex,
	but using RWMutex in struct we can simply it
*/
//- start RWMutex section

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.Lock()
	balance := account.Balance
	account.RWMutex.Unlock()

	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				// write to variable
				account.AddBalance(1)

				// read from variable
				fmt.Println(account.GetBalance())
			}

		}()

	}

	time.Sleep(5 * time.Second)
}

//- end of RWMutex

/*
## Pool
	is implementation of object pool pattern that store data, then to use it we can take from the pool, and store it again to pool.
	it will get data randomly from pool using Get, and store it again using Put.
*/

func TestPoolObjectPattern(t *testing.T) {
	// create pool with default value
	pool := sync.Pool{
		New: func() interface{} {
			return "Default-Value" // if poll is empty the default value will this
		},
	}

	pool.Put("Satrya")
	pool.Put("Budi")
	pool.Put("Pratama")

	for i := 0; i < 20; i++ {
		go func() {
			data := pool.Get() // this random get from pool

			fmt.Println(data)
			time.Sleep(1 * time.Second) // for latency

			pool.Put(data) // put it again
		}()
	}

	time.Sleep(4 * time.Second)
}

/*

 */
