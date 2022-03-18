package main

import "fmt"

func main() {

	ch := make(chan int)
	// processor in the middle
	outCh := proceed(ch)

	// producer
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("produce: %d\n", i)
			ch <- i
		}
		close(ch)
	}()

	// consumer
	quit := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			data := <-outCh
			fmt.Printf("consume: %d\n", data)
		}
		quit <- true
	}()

	// wait
	<-quit
}

// param : receieved only channel
// return : send only channel
func proceed(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		// keep consuming in channel until it's closed.
		for i := range in {
			out <- i * 2
		}
		close(out)
	}()
	return out
}

/* Result - producer -> process -> consumer, each run by different goroutines

produce: 0
produce: 1
produce: 2
consume: 0
consume: 2
consume: 4
produce: 3
produce: 4
produce: 5
consume: 6
consume: 8
consume: 10
produce: 6
produce: 7
produce: 8
consume: 12
consume: 14
consume: 16
produce: 9
consume: 18
*/
