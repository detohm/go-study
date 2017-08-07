package main

import "fmt"

func main() {

	fmt.Println("1")

	defer fmt.Println("7")
	defer fmt.Println("6")

	fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")

	defer fmt.Println("5")
}
