package main

import (
	"fmt"
)

func main() {
	fmt.Println("Lesson 8: Channels")

	/*
		Channels in Go are a fundamental concurrency primitive that allow goroutines (Go's lightweight threads) to communicate with each other and synchronize their execution. They provide a way for one goroutine to send values to another goroutine.
	*/

	// creates an unbuffered channel.
	var c = make(chan int)

	fmt.Println(c)

	// attempts to send the value 1 into the channel.
	// c <- 1

	// attempts to receive a value from the channel.
	// var i = <-c

	/*
		This will cause a "fatal error: all goroutines are asleep - deadlock!"
		The key issue is that unbuffered channels in Go require synchronization between the sender and receiver. When you try to send a value on an unbuffered channel (c <- 1), the send operation will block until there's a corresponding receive operation ready to accept the value.
		In your code, the send operation (c <- 1) blocks, waiting for a receiver. However, the program can't proceed to the receive operation (<- c) because it's blocked on the send. This results in a deadlock - the program can't make progress because both operations are waiting for each other.
	*/

	// To fix this, we have to use a goroutine for either the send or receive operation:
	// Update the channel using a goroutine
	go process(c)

	// Once the goroutine is done and c has been updated, print it.
	fmt.Println(<-c)

	// What if we want to loop over a chan an update it?
	var cTwo = make(chan int)

	// starts a new goroutine that runs the processTwo function, passing the channel cTwo as an argument.
	go processTwo(cTwo)

	//loop iterates over values received from the channel. This loop will continue until the channel is closed.
	for i := range cTwo {

		// prints each value received from the channel.
		fmt.Println(i)
	}
}

func process(c chan int) {
	c <- 123
}

func processTwo(c chan int) {
	// The defer keyword in Go schedules a function call (in this case, close(c)) to be run immediately before the function returns.
	defer close(c)
	for i := 0; i <= 5; i++ {
		c <- i
	}
}
