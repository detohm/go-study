package main

import "fmt"

func main() {

	var slice []int
	recentCap := cap(slice)

	for i := 1; i < 5000; i++ {
		slice = append(slice, i)

		if recentCap != cap(slice) {
			fmt.Printf("header-addr:%p arr-addr:%p len:%d cap:%d\n",
				&slice,
				&slice[0],
				len(slice),
				cap(slice))
			recentCap = cap(slice)
		}
	}
}

/* Result

header-addr:0xc0000a4018 arr-addr:0xc0000b2008 len:1 cap:1
header-addr:0xc0000a4018 arr-addr:0xc0000b2020 len:2 cap:2
header-addr:0xc0000a4018 arr-addr:0xc0000b8020 len:3 cap:4
header-addr:0xc0000a4018 arr-addr:0xc0000aa080 len:5 cap:8
header-addr:0xc0000a4018 arr-addr:0xc0000ba000 len:9 cap:16
header-addr:0xc0000a4018 arr-addr:0xc0000bc000 len:17 cap:32
header-addr:0xc0000a4018 arr-addr:0xc0000be000 len:33 cap:64
header-addr:0xc0000a4018 arr-addr:0xc0000c0000 len:65 cap:128
header-addr:0xc0000a4018 arr-addr:0xc0000c2000 len:129 cap:256
header-addr:0xc0000a4018 arr-addr:0xc0000c4000 len:257 cap:512
header-addr:0xc0000a4018 arr-addr:0xc0000c6000 len:513 cap:1024
header-addr:0xc0000a4018 arr-addr:0xc0000c8000 len:1025 cap:1280
header-addr:0xc0000a4018 arr-addr:0xc0000d2000 len:1281 cap:1696
header-addr:0xc0000a4018 arr-addr:0xc0000dc000 len:1697 cap:2304
header-addr:0xc0000a4018 arr-addr:0xc0000ee000 len:2305 cap:3072
header-addr:0xc0000a4018 arr-addr:0xc0000f4000 len:3073 cap:4096
header-addr:0xc0000a4018 arr-addr:0xc000100000 len:4097 cap:5120

Note: As keeping append element into slice, cap will be adjusted automatically.
Per observation, Cap will be increase doubly until 1024.
After that it will increase but not double manner.

As you can see, once reaching the cap, it allocated the new bigger array
(the arr-addr is changing) and all elements are copied into the new array.
However, the header is still at the same address.

*/
