package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			fmt.Println("print from go routine")
		}()
	}
	wg.Wait()
}

/* Result
print from go routine
print from go routine
print from go routine
print from go routine
print from go routine
*/
