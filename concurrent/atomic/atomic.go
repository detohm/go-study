package main

import (
	"fmt"
	"sync"
	"sync/atomic"
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

You can use command go build -race concurrent/atomic/atomic.go
for race detection
*/

func fixedRace() {
	var count int64 // change from int to int64
	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&count, 1) // apply atomic operation
				// innerCount := count
				// innerCount++
				// time.Sleep(1 * time.Microsecond)
				// count = innerCount

			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("count with atomic: %d\n", count)
}

/* Result
count with mutex: 500
The result is expected due to using atomic operation which synchronisation
at hardware level.
Please notice that we need to change type of variable `count` from int to int64
*/
