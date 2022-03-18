package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx1 := context.Background()
	ctx2, cancel2 := context.WithCancel(ctx1)
	ctx3, cancel3 := context.WithCancel(ctx2)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("cancelling ctx2")
		cancel2()
	}()

	for {
		select {
		case <-ctx2.Done():
			fmt.Println("ctx2 done")
		case <-ctx3.Done():
			fmt.Println("ctx3 done")
			cancel3() // should not affected as it's already cancelled by propagation
			return
		}
	}

}

/* Result
ctx3 will got cancelled as its parent (ctx2) got cancelled.
for-select loop will repeatly select ctx2.Done as it's closed channel
until ctx3 send the close signal and hit return statement

Output:
cancelling ctx2
ctx2 done
ctx2 done
ctx2 done
ctx3 done
*/
