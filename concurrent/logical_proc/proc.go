package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	runtime.GOMAXPROCS(1) // the number is map 1:1 to OS Thread
	proceed()

	// runtime.GOMAXPROCS(4) // try this with comment out the above
	// proceed()
}

func proceed() {

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		for i := 0; i < 50; i++ {
			fmt.Printf("%d ", i)
		}
		wg.Done()
	}()

	go func() {
		for i := 'a'; i < 'z'; i++ {
			fmt.Printf("%c ", i)
		}
		wg.Done()
	}()

	go func() {
		for i := 1000; i < 1050; i++ {
			fmt.Printf("%d ", i)
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println()
}

/* Result
as specified GOMAXPROCS=4, result is scatter due to multiple os threads
if you have more than 1 core processor or 1 core with hyper threading
you will get benefit from parallelism

1000 1001 1002 1003 1004 1005 1006 1007 a b c d e f g h i j k l m n o p
q r s t u v w x y 1008 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19
20 21 22 1009 1010 1011 1012 1013 1014 1015 1016 1017 1018 1019 1020 1021
1022 1023 1024 1025 1026 1027 1028 1029 1030 1031 1032 1033 1034 1035 1036
1037 1038 1039 1040 1041 1042 1043 1044 1045 1046 1047 1048 1049 23 24 25
26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49

Compare to the result from GOMAXPROCS=1. It is more in-order
as it run in single OS thread. Even though specified to use only
one single OS thread, you can spawn many of go routine as it's considered
as a light weight thread. Go runtime will take care all the scheduling process
of multiple go routines in very optimised manner.

1000 1001 1002 1003 1004 1005 1006 1007 1008 1009 1010 1011 1012 1013 1014
1015 1016 1017 1018 1019 1020 1021 1022 1023 1024 1025 1026 1027 1028 1029
1030 1031 1032 1033 1034 1035 1036 1037 1038 1039 1040 1041 1042 1043 1044
1045 1046 1047 1048 1049 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19
20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44
45 46 47 48 49 a b c d e f g h i j k l m n o p q r s t u v w x y
*/
