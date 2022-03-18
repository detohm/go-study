package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx2, cancel := context.WithTimeout(ctx, 2*time.Second)

	for {
		select {
		case <-ctx2.Done(): // blocked until receive done signal
			fmt.Println("timeout signal caught!")
			cancel()
			return
		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Println("waiting...")
		}
	}
}

/* Result - ctx2 send the done signal after 2 seconds

Output:
waiting...
waiting...
waiting...
waiting...
timeout signal caught!
*/
