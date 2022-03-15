package main

import (
	"fmt"
	"time"
)

func main() {
	go work()
	time.Sleep(2 * time.Second)
	fmt.Println("print from main")
}

func work() {
	fmt.Println("print from another go routine")
}

/* Result (most of the time)
print from another go routine
print from main

Note: main routine sleeps for 2 seconds
so, most of the time, 'work' routine finished before main routine
*/
