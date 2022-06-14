package main

import (
	"fmt"
	"os"
)

func main() {
	// bytes, err := ioutil.ReadFile("./created-file.txt")
	bytes, err := os.ReadFile("./created-file.txt") // go 1.6+
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s", bytes)
}
