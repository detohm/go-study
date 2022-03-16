package main

import "sync"

func main() {

	var wg sync.WaitGroup

	ch1 := make(chan int)
	ch2 := make(chan int)

	wg.Add(2)

	go func() {
		fromCh2 := <-ch2
		ch1 <- fromCh2
		wg.Done()
	}()

	go func() {
		fromCh1 := <-ch1
		ch2 <- fromCh1
		wg.Done()
	}()

	wg.Wait()
}

/* Result - each go routine are waiting for each other channel, deadlock occurs
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
sync.runtime_Semacquire(0x0)
        /usr/local/Cellar/go/1.17.8/libexec/src/runtime/sema.go:56 +0x25
sync.(*WaitGroup).Wait(0x0)
        /usr/local/Cellar/go/1.17.8/libexec/src/sync/waitgroup.go:130 +0x71
main.main()
        go-study/concurrent/problem/deadlock/channel/channel_deadlock.go:26 +0x149

goroutine 17 [chan receive]:
main.main.func1()
        go-study/concurrent/problem/deadlock/channel/channel_deadlock.go:15 +0x3d
created by main.main
        go-study/concurrent/problem/deadlock/channel/channel_deadlock.go:14 +0xcf

goroutine 18 [chan receive]:
main.main.func2()
        go-study/concurrent/problem/deadlock/channel/channel_deadlock.go:21 +0x3d
created by main.main
        go-study/concurrent/problem/deadlock/channel/channel_deadlock.go:20 +0x13f
exit status 2
*/
