package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := gen(1, 10)
	ch2 := gen(11, 20)

	// fan-in pattern: aggregating multiple input channel to
	// one output channel
	chFanIn := make(chan int)
	go func() {
		for {
			select {
			case n1 := <-ch1:
				chFanIn <- n1
			case n2 := <-ch2:
				chFanIn <- n2
			}
		}
	}()

	// long-run loop with receiving output from fan-in
	for {
		select {
		case fin := <-chFanIn:
			fmt.Println(fin)
		default:
			time.Sleep(time.Second)
		}
	}

}

// data generator function
func gen(start int, end int) <-chan int {

	// out channel
	ch := make(chan int)

	go func() {
		for i := start; i < end; i++ {
			ch <- i
			time.Sleep(200 * time.Millisecond)
		}
	}()

	return ch
}

/* Result
1
2
11
12
13
3
14
4
15
5
16
6
17
7
18
8
9
19
*/
