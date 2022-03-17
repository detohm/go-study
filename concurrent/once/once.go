package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	var once sync.Once
	for i := 0; i < 10; i++ {
		go func(ii int) {
			once.Do(func() {
				fmt.Printf("I'm called from %d\n", ii)
			})
			wg.Done()
		}(i)
	}
	wg.Wait()
}

/* Result - the func() inside once.Do is called only one time

I'm called from 0
*/
