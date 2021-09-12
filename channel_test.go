package goroutine

import (
	"fmt"
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

func ChannelParam(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Hello World2"
}

func TestChannelasParam(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go ChannelParam(channel)
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
