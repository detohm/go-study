package main

import "fmt"

func main() {

	// --- int ---
	sl := []int{1, 2, 3, 4}

	sl[1] = 100
	fmt.Printf("sl: %v\n", sl)
	// 1 100 3 4

	back := sl[len(sl)-1]
	back = back + 2
	fmt.Printf("sl: %v\n", sl)
	// 1 100 3 4 - no side effect to slice

	back2 := &sl[len(sl)-1]
	*back2 = *back2 + 2
	fmt.Printf("sl: %v\n", sl)
	// 1 100 3 6 - side effect to slice

	sl = sl[:len(sl)-1]
	fmt.Printf("sl: %v\n", sl)
	// 1 100 3

	// --- string ---
	sl2 := []string{"a", "b", "c", "d"}

	sl2[1] = "bbb"
	fmt.Printf("sl2: %v\n", sl2)
	// a bbb c d

	back3 := sl2[len(sl2)-1]
	back3 = back3 + "dd"
	fmt.Printf("sl2: %v\n", sl2)
	// a bbb c d - no side effect

	back4 := &sl2[len(sl2)-1]
	*back4 = *back4 + "dd"
	fmt.Printf("sl2: %v\n", sl2)
	// a bbb c ddd - side effect to slice

	// pointer to the last element in sl2 slice
	back5 := &sl2[len(sl2)-1]

	// reslice sl2 as index 0 to len - 2
	// you cannot access the last one(ddd) from sl2 anymore
	sl2 = sl2[:len(sl2)-1]
	fmt.Printf("sl2: %v\n", sl2)

	// since back5 is still pointed to ddd
	// it's still accessible even though it is not in slice anymore
	fmt.Printf("back5: %v\n", *back5)

	// This is very interesting!!!
	// Try to append the sl2 with new element
	sl2 = append(sl2, "eee")

	// back5 refer to the new element!!!
	// as back5 is refer to the memory address
	// next to the current slice
	// very interesting property !!!
	fmt.Printf("back5: %v\n", *back5)
}
