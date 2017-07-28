package main

import "fmt"

func main() {

	var s = make([]int, 3)

	fmt.Println(s)

	s[0] = 20
	s[1] = 32
	s[2] = 34

	fmt.Println(s)

	var t = s[0:0]
	fmt.Println(t)

	t = s[0:1]
	fmt.Println(t)

	t = s[0:2]
	fmt.Println(t)

	t = s[0:3]
	fmt.Println(t)

	t = s[0:]
	fmt.Println(t)

	t = s[:1]
	fmt.Println(t)

	t = s[1:2]
	fmt.Println(t)

	t = s[2:3]
	fmt.Println(t)

	t = s[2:2]
	fmt.Println(t)

}
