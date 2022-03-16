package main

import (
	"fmt"
)

func main() {
	channel()
}

func channel() {
	ch := make(chan int)
	go func() {
		//time.Sleep(300 * time.Millisecond)
		fmt.Println("sending 5 from go routine")
		ch <- 5
	}()
	res := <-ch
	fmt.Printf("received %d\n", res)
}

/* Result always
sending 5 from go routine
received 5

This is synchronous behavior as the main routine waits result from channel
at res := <-ch line
*/
