package main

import "fmt"

func main() {
	ch := make(chan int, 2) // buffered channel
	ch <- 1
	ch <- 2
	ch <- 3
	for i := 0; i < 2; i++ {
		fmt.Printf("%d", <-ch)
	}
}

/* Result - send data more than cap of buffered channel,
sender will got blocked.

fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
        go-study/concurrent/problem/deadlock/buffered_channel/buffered_channel.go:9 +0x5c
exit status 2
*/
