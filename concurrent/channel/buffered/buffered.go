package main

import (
	"bytes"
	"fmt"
	"os"
)

// this is modified from the `range` example (../range/range.go)
func main() {
	var b bytes.Buffer

	// modified from ch := make(chan int)
	ch := make(chan int, 3)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Fprintf(&b, "Sending:%d\n", i)
			ch <- i
		}
		close(ch) // if you forget to close, it will be deadlock occurs
	}()

	for in := range ch { // keep waiting from channel until it's closed
		fmt.Fprintf(&b, "Received:%d\n", in)
	}

	// use buffer instead of print directly to stdout
	// in order to demonstate behavior without interfering from IO task
	b.WriteTo(os.Stdout)
}

/* result - less scatter comparing to the unbuffered channel in range example
as there is no waiting in sender until channel is full
Sending:0
Sending:1
Sending:2
Sending:3
Sending:4
Received:0
Received:1
Received:2
Received:3
Received:4
Sending:5
Sending:6
Sending:7
Sending:8
Sending:9
Received:5
Received:6
Received:7
Received:8
Received:9
*/
