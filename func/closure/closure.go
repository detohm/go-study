package main

import "fmt"

func counter() func() int {
	i := 10
	return func() int {
		i--
		return i
	}
}

func main() {

	tick := counter()

	fmt.Println(tick()) //9
	fmt.Println(tick()) //8
	fmt.Println(tick()) //7
	fmt.Println(tick()) //6
	fmt.Println(tick()) //5

}
