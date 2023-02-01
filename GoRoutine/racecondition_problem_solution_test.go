package GoRoutine

import (
	"fmt"
	"sync"
	"sync/atomic"
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
## Sync Map
	Similiar with map, but it designed for goroutine process to handle race condition.
*/
// - start
func AddToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()
	group.Add(1)

	data.Store(value, value) // store key and value
}

func TestSyncMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go AddToMap(
			data, i, group) // add goroutine
	}

	// wait until process finished
	group.Wait()

	fmt.Println("Printed data")
	// iterate data using range method to print data
	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})
}

// --end

/*
## Sync.Cond
	Implements condition based locking. Need locker such as Mutex or RWMutex.
	There's signal() to continue process wait() one by one inside goroutine,
	also broadcast() to continoue process waat() for all.
*/
// -- start

func WaitSyncCondition(value int, cond *sync.Cond, group *sync.WaitGroup) {
	defer group.Done()
	group.Add(1)

	// implements lock
	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done", value)
	cond.L.Unlock()
}

func TestSyncCondition(t *testing.T) {
	lock := &sync.Mutex{}
	cond := sync.NewCond(lock)

	group := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go WaitSyncCondition(i, cond, group)
	}

	// send signal to running one by one
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal() // send signal to condition
		}
	}()

	// or send broadcast to all goroutine to skipped wait()
	go func() {
		time.Sleep(1 * time.Second)
		cond.Broadcast()
	}()

	group.Wait()
}

/*
## Atomic
	This package helps simply of not using Mutex or RWMutex in concurrent.
	And it applies in primitive variable like numeric but not struct.
*/
// - start

func TestAtomic(t *testing.T) {
	var value int64 = 0

	group := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		go func() {
			group.Add(1)

			for j := 0; j < 100; j++ {
				atomic.AddInt64(&value, 1)
			}

			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Value : ", value)
}

// - end
