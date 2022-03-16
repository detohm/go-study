package main

import (
	"fmt"
	"sync"
	"time"
)

type data struct {
	mu    sync.Mutex
	value int
}

func main() {
	data1 := data{sync.Mutex{}, 1}
	data2 := data{sync.Mutex{}, 2}

	var wg sync.WaitGroup
	wg.Add(2)
	go process(&data1, &data2, &wg)
	go process(&data2, &data1, &wg)
	wg.Wait()
}

func process(d1 *data, d2 *data, wg *sync.WaitGroup) {

	d1.mu.Lock()

	// if remove sleep, it will fix the issue most of the time. HOWEVER
	// removing this is not the good solution as it's non-deterministic behavior
	time.Sleep(time.Second)

	d2.mu.Lock()
	fmt.Printf("%d", d1.value+d2.value)
	d2.mu.Unlock()

	d1.mu.Unlock()
	wg.Done()
}

/* Result
Mutex lock are waiting for each other forever
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
sync.runtime_Semacquire(0x0)
        /usr/local/Cellar/go/1.17.8/libexec/src/runtime/sema.go:56 +0x25
sync.(*WaitGroup).Wait(0x0)
        /usr/local/Cellar/go/1.17.8/libexec/src/sync/waitgroup.go:130 +0x71
main.main()
        go-study/concurrent/problem/deadlock/mutex/mutex_deadlock.go:22 +0x159

goroutine 18 [semacquire]:
sync.runtime_SemacquireMutex(0xc00003e748, 0x8e, 0x88675c73a7ff)
        /usr/local/Cellar/go/1.17.8/libexec/src/runtime/sema.go:71 +0x25
sync.(*Mutex).lockSlow(0xc0000b2020)
        /usr/local/Cellar/go/1.17.8/libexec/src/sync/mutex.go:138 +0x165
sync.(*Mutex).Lock(...)
        /usr/local/Cellar/go/1.17.8/libexec/src/sync/mutex.go:81
main.process(0xc0000b2010, 0xc0000b2020, 0x0)
        go-study/concurrent/problem/deadlock/mutex/mutex_deadlock.go:31 +0x6f
created by main.main
        go-study/concurrent/problem/deadlock/mutex/mutex_deadlock.go:20 +0xd6

goroutine 19 [semacquire]:
sync.runtime_SemacquireMutex(0xc00003ef48, 0x8e, 0x88675c73611c)
        /usr/local/Cellar/go/1.17.8/libexec/src/runtime/sema.go:71 +0x25
sync.(*Mutex).lockSlow(0xc0000b2010)
        /usr/local/Cellar/go/1.17.8/libexec/src/sync/mutex.go:138 +0x165
sync.(*Mutex).Lock(...)
        /usr/local/Cellar/go/1.17.8/libexec/src/sync/mutex.go:81
main.process(0xc0000b2020, 0xc0000b2010, 0x0)
        go-study/concurrent/problem/deadlock/mutex/mutex_deadlock.go:31 +0x6f
created by main.main
        go-study/concurrent/problem/deadlock/mutex/mutex_deadlock.go:21 +0x14f
exit status 2
*/
