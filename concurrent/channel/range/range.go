package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var b bytes.Buffer

	ch := make(chan int)
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

/* result
Sending:0
Sending:1
Received:0
Received:1
Sending:2
Sending:3
Received:2
Received:3
Sending:4
Sending:5
Received:4
Received:5
Sending:6
Sending:7
Received:6
Received:7
Sending:8
Sending:9
Received:8
Received:9
*/
