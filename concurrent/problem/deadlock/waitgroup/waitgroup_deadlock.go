package main

import "sync"

func main() {
	var wg sync.WaitGroup
	wg.Add(6)
	for i := 0; i < 5; i++ {
		wg.Done()
	}
	wg.Wait()
}

/* Result - as workgroup will be wait forever, error occurs

fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
sync.runtime_Semacquire(0xc)
        /usr/local/Cellar/go/1.17.8/libexec/src/runtime/sema.go:56 +0x25
sync.(*WaitGroup).Wait(0x0)
        /usr/local/Cellar/go/1.17.8/libexec/src/sync/waitgroup.go:130 +0x71
main.main()
        go-study/concurrent/problem/deadlock/waitgroup/waitgroup_deadlock.go:11 +0x68
exit status 2
*/
