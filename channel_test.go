package goroutine

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Hello World"
		fmt.Println("Done sending data to channel")
	}()
	data := <-channel
	fmt.Println(data)

	time.Sleep(6 * time.Second)
}

/*
	<--------------------------------Channel Sebagai Parameter---------------------------------------->
*/

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Hello World2"
}

func TestChannelasParam(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)
	data := <-channel
	fmt.Println(data)

	time.Sleep(6 * time.Second)
}

/*
	<--------------------------------Channel In & Out---------------------------------------->
*/
//Untuk Channel yg hanya bisa menerima data maka tanda panah dibelakang keyword chan
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Hello World3"
}

//Untuk Channel yg hanya bisa mengirim data maka tanda panah didepan keyword chan
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOut(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

/*
	<--------------------------------Buffered Channel---------------------------------------->
*/

func TestBuffered(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Hello World"  // Received data 1
	channel <- "Hello World2" // Received data 2
	channel <- "hello world3" // Received data 3
	//channel <- "Hello World4" // Received data more than channel capacity will causes error

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("Done Sending data")
}

/*
	<--------------------------------Range Channel---------------------------------------->
*/

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		defer close(channel)
	}()

	for data := range channel {
		fmt.Println("Received data ", data)
	}
	fmt.Println("done sending data")
}

/*
	<--------------------------------Select Channel---------------------------------------->
*/

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0

	for {
		select {
		case data := <-channel1:
			fmt.Println("Data From Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Received From Channel 2", data)
			counter++
		default:
			fmt.Println("Waiting...")
		}
		if counter == 2 {
			break
		}
	}
}

/*
	<--------------------------------Race Condition---------------------------------------->
*/

func TestRaceCondition(t *testing.T) {
	x := 0

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				x = x + 1
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Counter ", x)
}

/*
	<--------------------------------Sync.Mutex---------------------------------------->
*/

func TestSyncMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Counter ", x)
}

/*
	<--------------------------------Read Write Mutex---------------------------------------->
*/

type BankAccount struct {
	RwMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RwMutex.Lock()
	account.Balance = account.Balance + amount
	account.RwMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RwMutex.RLock() //Read Lock
	balance := account.Balance
	account.RwMutex.RUnlock() //Read Unlock
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Total = ", account.GetBalance())
}

/*
	<--------------------------------Dead Lock---------------------------------------->
*/

type UserBalance struct {
	sync.Mutex
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
	user1.Lock()
	fmt.Println("Lock", user1.Name)
	user1.Change(-amount)

	time.Sleep(2 * time.Second)

	user2.Lock()
	fmt.Println("Lock", user2.Name)
	user2.Change(amount)

	time.Sleep(2 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadLock(t *testing.T) {

}
