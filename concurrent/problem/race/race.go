package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	race()
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
A data race occur

You can use command go build -race concurrent/problem/race/race.go
then run ./main it will populate data rece detection such below,
==================
WARNING: DATA RACE
Read at 0x00c00013c018 by goroutine 8:
  main.race.func1()
      go-study/concurrent/problem/race/race.go:21 +0x46

Previous write at 0x00c00013c018 by goroutine 7:
  main.race.func1()
      go-study/concurrent/problem/race/race.go:24 +0x67

Goroutine 8 (running) created at:
  main.race()
      go-study/concurrent/problem/race/race.go:19 +0x8f
  main.main()
      go-study/concurrent/problem/race/race.go:10 +0x24

Goroutine 7 (running) created at:
  main.race()
      go-study/concurrent/problem/race/race.go:19 +0x8f
  main.main()
      go-study/concurrent/problem/race/race.go:10 +0x24
==================
count: 153
Found 1 data race(s)
*/
