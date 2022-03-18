package main

import (
	"fmt"
)

func main() {

	// simulate long running process
	for {
		// passing nil channel
		go task(nil)
	}
}

func task(in chan int) {
	input := <-in // this will wait forever
	if input > 0 {
		fmt.Println("greater than zero")
	}
}

/* Result
Memory consumption will be increasing until out of memory occurs
It's go routines leak as task in go routine cannot be completed
*/
