package main

import "fmt"

func main() {
	var a int = 5
	fmt.Printf("addr:%p value:%d\n", &a, a)

	b := []int{1, 2, 3, 4}
	fmt.Printf("b-addr:%p b-inner-arr-addr:%p b-value:%v\n", &b, &b[0], b)

	c := &b
	fmt.Printf("c-addr:%p c-point-to-addr:%p c-value:%v\n", &c, c, c)

	passByPointer(c)

	/* Result
	b-addr:0xc0000a4018 b-inner-arr-addr:0xc0000ba000 value:[1 2 3 4]
	c-addr:0xc0000ac020 c-point-to-addr:0xc0000a4018 c-value:&[1 2 3 4]
	inside function: d-addr:0xc0000ac028, d-point-to-addr:0xc0000a4018

	Note: Per observation, c and d are point to the same address (b-addr).
	When passing with pointer, the pointer address are changed
	(new allocation in argument)
	but the address value inside pointer are the same.
	*/

	fmt.Printf("b-value:%v\n", b)
	/* Result
	b-value:[1234 2 3 4]
	Note: value got updated inside passByPointer function
	*/
}

func passByPointer(d *[]int) {
	fmt.Printf("inside function: d-addr:%p, d-point-to-addr:%p\n", &d, d)
	(*d)[0] = 1234
}
