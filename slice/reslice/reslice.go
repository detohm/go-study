package main

import "fmt"

func main() {

	slice := []int{1, 2, 3}
	slice = append(slice, 4)
	slice = append(slice, 5)
	slice = append(slice, 6)

	fmt.Printf("slice - header-addr:%p arr-addr:%p len:%d cap:%d value:%v\n",
		&slice,
		&slice[0],
		len(slice), cap(slice),
		slice)
	/* Result
	slice - header-addr:0xc00000c030 arr-addr:0xc00001c180
	len:6 cap:6
	value:[1 2 3 4 5 6]
	*/

	slice2 := slice[0:2]
	fmt.Printf("slice2 - header-addr:%p arr-addr:%p len:%d cap:%d value:%v\n",
		&slice2,
		&slice2[0],
		len(slice2), cap(slice2),
		slice2)
	/* Result
	slice2 - header-addr:0xc00000c060 arr-addr:0xc00001c180
	len:2 cap:6
	value:[1 2]
	*/

	slice3 := slice[1:2]
	fmt.Printf("slice3 - header-addr:%p arr-addr:%p len:%d cap:%d value:%v\n",
		&slice3,
		&slice3[0],
		len(slice3), cap(slice3),
		slice3)
	/* Result
	slice3 - header-addr:0xc00000c090 arr-addr:0xc00001c188
	len:1 cap:5
	value:[2]
	*/

	fmt.Println()

	slice[0] = 1234
	fmt.Printf("slice[0]:%d slice2[0]:%d\n", slice[0], slice2[0])
	/* Result
	slice[0]:1234 slice2[0]:1234
	Note: if changing the first value in `slice`,
	the first value of `slice2` also be affected.
	*/

	// append many elements to force `slice` to re-allocation with bigger array
	for i := 0; i < 10000; i++ {
		slice = append(slice, i)
	}

	slice[0] = 5678
	fmt.Printf("slice[0]:%d slice2[0]:%d\n", slice[0], slice2[0])
	/* Result
	slice[0]:5678 slice2[0]:1234
	Note: interesting property! try changing the first value in `slice` again,
	the first value of `slice2` is not affected.
	This behavior is different from the first one that we change value to 1234
	*/

	fmt.Printf("slice - header-addr:%p arr-addr:%p\n",
		&slice,
		&slice[0])
	fmt.Printf("slice2 - header-addr:%p arr-addr:%p\n",
		&slice2,
		&slice2[0])

	/* Result
	slice - header-addr:0xc0000a4018 arr-addr:0xc000120000
	slice2 - header-addr:0xc0000a4048 arr-addr:0xc0000b4030
	Note: underlying array is different because of dynamic allocation !!!
	*/
}
