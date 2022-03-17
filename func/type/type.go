package main

import "fmt"

type processFunc func(int)

func main() {
	var fc processFunc
	fc = func(i int) {
		fmt.Println(i)
	}
	fc(1)
}
