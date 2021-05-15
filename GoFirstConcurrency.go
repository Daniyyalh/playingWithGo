package main

import (
	"fmt"
	"math/rand"
	"time"
)

func helloWorld(num int) {
	// Note that in the loop there's no parantheses, unlike java. Otherwise the structure of the loop is the same
	for i := 0; i < 10; i++ {
		fmt.Printf("From thread %d\n", num)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

// Takes a channel and adds "ping" to it
func pinger(c chan string) {
	// a loop that doesn't terminate
	for i := 0; ; i++ {
		c <- "ping"
	}
}

// Takes a channel that can only be sent to. Note the argument. It can't recieve anything
func pinger2(c chan<- string) {
	// a loop that doesn't terminate
	for i := 0; ; i++ {
		c <- "ping"
	}
}

// Takes a channel and adds "pong" to it
func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

// prints the message in the channel and sleeps
func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

// // Takes a channel that can only recieve. Note the argument. It can't send anything
func printer2(c <-chan string) {

	for {
		// can't do c <- "pong" because that's recieving
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {

	// example 1
	// The block below is used to show how to make things asynchronous

	// for i := 0; i < 10; i++ {
	// 	go helloWorld(i)
	// }

	// // Needed this ask for an input because otherwise the program finishes before the other functions are called
	// var input string
	// fmt.Scanln(&input)

	// example 2, with channels
	// since we didn't pass in any argument to `make`, the channel will NOT work asynchronously

	// var c chan string = make(chan string)

	// go pinger(c)
	// go ponger(c)
	// go printer(c)

	// var input string
	// fmt.Scanln(&input)

	// example 3, with channels
	// since we passed in an argument to `make`, the channel WILL work asynchronously

	// var c chan string = make(chan string, 2)

	// go pinger(c)
	// go ponger(c)
	// go printer(c)

	// var input string
	// fmt.Scanln(&input)

	// example 4, with `select`. select let's us do specific things each time something is written to a channel

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println("Message 1", msg1)
		case msg2 := <-c2:
			fmt.Println("Message 2", msg2)
		case <-time.After(time.Second):
			fmt.Println("timeout")
		}
	}

}
