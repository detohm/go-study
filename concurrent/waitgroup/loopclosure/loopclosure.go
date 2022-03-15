package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("wrong example")
	wrong()
	fmt.Println("correct example")
	right()
}

func wrong() {
	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			fmt.Printf("print: %d\n", i)
		}()
	}
	wg.Wait()
}

/* Wrong Result
print: 5
print: 3
print: 5
print: 5
print: 5

Note: As you can see the result, most of the result are 5 (last value) why?
because loop finishes running before any go routine actually start
and all of them refers to the same i (not a copy of it)
*/

func right() {
	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(ii int) {
			defer wg.Done()
			fmt.Printf("print: %d\n", ii)
		}(i)
	}
	wg.Wait()
}

/* Right Result
print: 4
print: 2
print: 3
print: 0
print: 1

Note: Fixed by passing i into go routine, make it captured by value
*/
