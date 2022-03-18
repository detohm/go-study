package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 10
		time.Sleep(2 * time.Second)
		close(ch1)
	}()
	count := 0
	for {
		select {
		case value, ok := <-ch1:
			fmt.Printf("received v:%d ok:%v\n", value, ok)
			count++
			if count > 10 {
				ch1 <- 1000
			}
		}
	}

}

/* Result
This example demonstrates 3 interesting behavior
1. after close the channel, receiver will got another signal (0,false).
2. if try to receive value from the close channel,
   it will not block and get (0, false) everytime.
2. if try to send value to closed channel, it will cause panic.

Output:
received v:10 ok:true
received v:0 ok:false
received v:0 ok:false
received v:0 ok:false
received v:0 ok:false
received v:0 ok:false
received v:0 ok:false
received v:0 ok:false
received v:0 ok:false
received v:0 ok:false
received v:0 ok:false
panic: send on closed channel

goroutine 1 [running]:
main.main()
        go-study/concurrent/channel/close/close.go:23 +0x145
exit status 2
*/
