package main

import "fmt"

func main() {

	// use anonymous function
	operate(func(a int, b int) int {
		return a + b
	}, 1, 2)

	// initialise function variable
	multiply := func(a int, b int) int {
		return a * b
	}
	operate(multiply, 4, 5)

	// use type declaration
	type opFunc func(int, int) int
	// declare without initialization
	var minus opFunc
	// initialise the function
	minus = func(a int, b int) int {
		return a - b
	}
	operate(minus, 10, 3)
}

func operate(f func(int, int) int, a int, b int) int {
	res := f(a, b)
	fmt.Println(res)
	return res
}
