package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Create a mutex to lock and unlock memory for Goroutines
var m = sync.RWMutex{}
// Create a wait group so we can tell go to wait for everything to finish before ending the main func
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	fmt.Println("Goroutines")

	/*
		Goroutines are one of the most powerful and distinctive features of Go. They allow you to execute functions concurrently, making it easier to build efficient, scalable, and concurrent programs.

		- Lightweight threads: Goroutines are much lighter than traditional operating system threads. They have very small initial memory overhead and can scale to hundreds of thousands of goroutines without consuming significant system resources.
		- Concurrency in Go: By launching goroutines, you can perform tasks concurrently, improving the responsiveness of your application. This is particularly useful in I/O-bound or CPU-bound operations.
		- Managed by the Go runtime: Goroutines are scheduled and managed by Go's runtime, which handles their execution across available system threads.
	*/

	t0 := time.Now()
	
	// if we didn't have the go keyword before the dbCall, it would take 2 seconds for each call synchronously. Long.
	// However, the main func will end before the dbCalls are returned. So we need to wait. We can do this using a wait group (Promise.all())
	for i := 0; i < len(dbData); i++ {
		// This just adds a counter to the wait group. It is removed when the wg.Done is called within the dbCall func
		wg.Add(1)
		go dbCall(i)
	}

	// This will wait until the wg count is down to 0
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\nThe results are %v", results)
}

func dbCall(i int) {
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is:", dbData[i])
	save(dbData[i])
	log()
	wg.Done()
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	// Sometimes we don't want a full lock, we just want to read the data, hence why we can use a read write mutex to achieve this.
	m.RLock()
	fmt.Printf("\nThe current results are: %v", results)
	m.RUnlock()
}