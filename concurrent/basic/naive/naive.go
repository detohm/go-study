package main

import "fmt"

func main() {
	go work()
	fmt.Println("print from main")
}

func work() {
	fmt.Println("print from another go routine")
}

/* Result (most of the time)
print from main

Note: main routine finished before another go routine to complete
*/
