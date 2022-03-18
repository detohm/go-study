package main

import (
	"context"
)

func main() {
	ctx := context.Background()

	// contextWithCancel contains its children context pointers
	// for cancel propagation
	ctx2, cancel := context.WithCancel(ctx)
	defer cancel()

	// 1. Fixed by using background context instead
	// ctx2 = context.Background()

	// big chunk of data to accelerate memory consumption
	str := "asdfasdfahsdjfhasdlkfhasdjfhalsdfhjasdflahsfjhasdlfhaksdjfhjaksfhakshdfjkdfhalsdkfhasdjlkfhalksdhfkaljsdfhlkajsdhflaskdfhalksdfkajshdfjalskdfashdfahsjkdf"
	count := 1

	// simulate long-run process like web server
	for {
		// time.Sleep(time.Nanosecond)
		ctx3, cancelCtx3 := context.WithCancel(ctx2)
		ctx3 = context.WithValue(ctx3, "key", count)
		ctx3 = context.WithValue(ctx3, "str", str)
		count++
		// pass contextwithcancel to ro routine
		go proceed(ctx3, cancelCtx3)
	}
}

func proceed(rCtx context.Context, cancel context.CancelFunc) {

	// 2. Fixed by cancel the child context after use
	// defer cancel()

	i := rCtx.Value("key").(int)

	if i < 0 {
		// this code never reach as i always greater than 0
		cancel()
	}
}

/* Result
Memory usage will keep increasing until out of memory occurs (mem leak)
as ctx2 is not blank context (background) but it's a contextWithCancel
which contains children context pointers. Therefore, GC cannot clean up
the child context as it's refered from their parent.

You can fix with 2 approaches below,
1. Change ctx2 to background context so that there is no children reference.
GC can clean up the child context.
2. defer cancel of child context after finished the job.
*/
