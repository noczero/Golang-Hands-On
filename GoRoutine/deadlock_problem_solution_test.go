package GoRoutine

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	Mutex   sync.Mutex // or direct sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	/*
	 Balance in User1 will debit and user2 will credit.

	 Deadlock Flow : user 1 lock, manipulate amount, then sleep to simulate latency without unlock,
	 user 2 lock, manipulate amount, then sleep to simulate latency without unlock,
	 finally unlock user1, and unlock user2

	 Solution :
	 to solve this every lock() should unlock() before continue to others process,

	*/
	user1.Lock()
	fmt.Println("Lock user1", user1.Name)
	user1.Change(-amount) // debit
	time.Sleep(1 * time.Second)

	// unlock should place to this, to prevent the deadlock by multiple goroutine.
	user1.Unlock()

	user2.Lock()
	fmt.Println("Lock user2", user1.Name)
	user2.Change(amount) // credit
	time.Sleep(1 * time.Second)

	//user1.Unlock() // this unlock in the end make deadlock when called by multiple goroutine.
	user2.Unlock()
}

func TestDeadLockDemo(t *testing.T) {
	user1 := UserBalance{
		Mutex:   sync.Mutex{},
		Name:    "Satrya",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Mutex:   sync.Mutex{},
		Name:    "Budi",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000) // transfer 100k from user1 to user2 using goroutine

	go Transfer(&user2, &user1, 200000) // transfer 200k from user2 to user1 using goroutine

	time.Sleep(3 * time.Second)

	fmt.Println("User 1 : ", user1.Name, "Balance : ", user1.Balance) // true balance will 1100k
	fmt.Println("User 2 : ", user2.Name, "Balance : ", user2.Balance) // true balance will 900k

	assert.Equal(t, 1100000, user1.Balance)
	assert.Equal(t, 900000, user2.Balance)
}
