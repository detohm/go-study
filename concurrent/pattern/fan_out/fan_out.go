package main

import (
	"fmt"
	"time"
)

func main() {
	ch := gen(0, 10)
	wch1 := worker(ch, 1)
	wch2 := worker(ch, 2)
	for {
		select {
		case i := <-wch1:
			fmt.Println(i)
		case j := <-wch2:
			fmt.Println(j)
		default:
			time.Sleep(time.Millisecond)
		}

	}
}

func gen(begin int, end int) <-chan int {
	out := make(chan int)
	go func() {
		for i := begin; i < end; i++ {
			out <- i
			time.Sleep(200 * time.Millisecond)
		}

	}()
	return out
}

// in is received only channel
// out is received only channel
func worker(in <-chan int, num int) <-chan int {
	out := make(chan int)
	go func() {
		for i := range in {
			fmt.Printf("processing from worker %d\n", num)
			out <- i * i
		}

	}()
	return out
}

/* Result
process from worker 1
0
process from worker 2
1
process from worker 1
4
process from worker 2
9
process from worker 1
16
process from worker 2
25
process from worker 1
36
process from worker 2
49
process from worker 1
64
process from worker 2
81
*/
