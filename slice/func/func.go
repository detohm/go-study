package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4}

	fmt.Printf("before: header-addr:%p arr-addr:%p [0]:%d len:%d cap:%d\n",
		&slice,
		&slice[0],
		slice[0],
		len(slice),
		cap(slice))

	change(slice)

	fmt.Printf("after: header-addr:%p arr-addr:%p [0]:%d len:%d cap:%d\n\n",
		&slice,
		&slice[0],
		slice[0],
		len(slice),
		cap(slice))
	/* Result
	before: header-addr:0xc00000c030 arr-addr:0xc00001a020 [0]:1 len:4 cap:4
	in func change: header-addr:0xc00000c048 arr-addr:0xc00001a020 [0]:11111 len:4 cap:4
	after: header-addr:0xc00000c030 arr-addr:0xc00001a020 [0]:11111 len:4 cap:4

	Note: change applied inside function also affects the caller
	*/

	change2(slice)

	fmt.Printf("after: header-addr:%p arr-addr:%p [0]:%d len:%d cap:%d\n\n",
		&slice,
		&slice[0],
		slice[0],
		len(slice),
		cap(slice))
	/* Result
	in func change2: header-addr:0xc00000c060 arr-addr:0xc000168000 [0]:22222 len:9999 cap:12288
	after: header-addr:0xc00000c030 arr-addr:0xc00001a020 [0]:11111 len:4 cap:4

	Note: change applied inside function not affects totally to the caller's slice

	interesting non-deterministic behavior (between calling func1 and func2)!!!
	this behavior happens because of dynamic allocation.
	*/

}

func change(sl []int) {

	sl[0] = 11111

	fmt.Printf("in func change: header-addr:%p arr-addr:%p [0]:%d len:%d cap:%d\n",
		&sl,
		&sl[0],
		sl[0],
		len(sl),
		cap(sl))
}

func change2(sl []int) {
	for i := 5; i < 10000; i++ {
		sl = append(sl, i)
	}
	sl[0] = 22222
	fmt.Printf("in func change2: header-addr:%p arr-addr:%p [0]:%d len:%d cap:%d\n",
		&sl,
		&sl[0],
		sl[0],
		len(sl),
		cap(sl))
}
