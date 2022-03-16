package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	race()
	fixedRace()
}

func race() {
	var count int
	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				innerCount := count
				innerCount++
				time.Sleep(1 * time.Microsecond)
				count = innerCount
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("count: %d\n", count)
}

/* Result
count: 102
why it's not 500?
This is not syncronised access the count variable
A data race occurs

You can use command go build -race concurrent/mutex/mutex.go
for race detection
*/

func fixedRace() {
	var count int
	var wg sync.WaitGroup
	wg.Add(5)

	var mu sync.Mutex

	for i := 0; i < 5; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mu.Lock() // start critical section
				innerCount := count
				innerCount++
				// time.Sleep(1 * time.Microsecond)
				count = innerCount
				mu.Unlock() // end critical section
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("count with mutex: %d\n", count)
}

/* Result
count with mutex: 500
The result is expected due to using mutex to lock the critical section
which synchronized access to the count variable.
*/
