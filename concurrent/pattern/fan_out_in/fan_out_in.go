package main

import (
	"fmt"
	"time"
)

func main() {
	genCh := gen(1)
	workerChs := []<-chan int{}
	for w := 1; w <= 5; w++ {
		workerChs = append(workerChs, worker(w, genCh))
	}

	aggCh := fanIn(workerChs)
	for {
		select {
		case msg := <-aggCh:
			fmt.Print(msg)
		}
	}

}

func gen(start int) chan int {
	ch := make(chan int)
	go func() {
		for i := start; i < 100; i++ {
			ch <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return ch
}

// fan out
func worker(name int, genCh <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		for {
			task := <-genCh
			fmt.Printf("process task %d on worker %d\n", task, name)
			ch <- task * task
		}
	}()
	return ch
}

// fan in as aggregation
func fanIn(channels []<-chan int) <-chan string {
	aggCh := make(chan string)
	for _, ch := range channels {
		go func(c <-chan int) {
			for v := range c {
				aggCh <- fmt.Sprintf("agg: %d\n", v)
			}
		}(ch)
	}
	return aggCh
}
