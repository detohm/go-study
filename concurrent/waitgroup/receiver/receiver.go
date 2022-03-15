package main

import (
	"fmt"
	"sync"
)

func main() {
	println("wrong")
	wrong()

	println()
	println("right")
	right()
}

type task int

func (t *task) print1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("print1: %d\n", *t)
}

func (t task) print2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("print2: %d\n", t)
}

func wrong() {
	tasks := []task{1, 2, 3, 4, 5}
	var wg sync.WaitGroup
	wg.Add(5)
	for _, v := range tasks {
		go v.print1(&wg)
	}
	wg.Wait()
}

func right() {
	tasks := []task{1, 2, 3, 4, 5}
	var wg sync.WaitGroup
	wg.Add(5)
	for _, v := range tasks {
		go v.print2(&wg)
	}
	wg.Wait()
}

/* Result
wrong
print1: 5
print1: 5
print1: 5
print1: 5
print1: 5

right
print2: 5
print2: 4
print2: 3
print2: 2
print2: 1

Note: very interesting behavior even though the code from calling side are
mostly identical (print1 and print2)

There are difference by using value and pointer receiver.
Value receiver avoid this problem as it's copy value to each
particular routime instead of referring to the same var.
*/
