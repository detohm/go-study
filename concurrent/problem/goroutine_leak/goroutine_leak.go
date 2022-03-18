package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	// --- non-related leak code
	quitObserver := make(chan bool)
	go observer(quitObserver)
	defer func() {
		quitObserver <- true
	}()
	// ---

	// simulate long running process
	for {
		// passing nil channel
		time.Sleep(time.Nanosecond)
		go task(nil)
	}
}

func task(in chan int) {
	input := <-in // this will wait forever
	if input > 0 {
		fmt.Println("greater than zero")
	}
}

// helper function for observing num of goroutines
func observer(quit chan bool) {
	for {
		select {
		case <-quit:
			return
		default:
			time.Sleep(time.Second)
			fmt.Printf("num of goroutines: %d\n", runtime.NumGoroutine())
		}
	}
}

/* Result
Memory consumption will be increasing until out of memory occurs
This is goroutines leak issue as a task in goroutine cannot be completed
due to blocked from nil channel.

The number of runtime goroutines will keep increasing
which affects memory consumption.

Output:
num of goroutines: 370824
num of goroutines: 761902
num of goroutines: 1266289
num of goroutines: 1573220
num of goroutines: 2067518
num of goroutines: 2347368
num of goroutines: 2694127
num of goroutines: 3122206
num of goroutines: 3561923
num of goroutines: 3943206
num of goroutines: 4289707
*/
