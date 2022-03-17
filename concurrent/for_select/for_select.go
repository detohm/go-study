package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	quitCh := make(chan bool)

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(500 * time.Millisecond)
			ch1 <- 1
		}
	}()
	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- 2
	}()
	go func() {
		time.Sleep(2 * time.Second)
		quitCh <- true
	}()

	for {
		select {
		case m1 := <-ch1:
			fmt.Printf("got %d from ch1\n", m1)
		case m2 := <-ch2:
			fmt.Printf("got %d from ch2\n", m2)
		case <-quitCh:
			fmt.Printf("got quite signal\n")
			return
			// default:
			// 	fmt.Printf("will not wait")
		}
	}

}

/* Result
got 1 from ch1
got 1 from ch1
got 1 from ch1
got quite signal

the for-select loop is alway waiting data for different channels
interesting that if we put default case, it will not wait anymore
*/
