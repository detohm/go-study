package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	data := []string{
		"This ",
		"is ",
		"a ",
		"book.",
	}
	for _, p := range data {
		n, err := os.Stdout.Write([]byte(p))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if n != len(p) {
			fmt.Println("failed to write data.")
			os.Exit(1)
		}

		time.Sleep(time.Second)
	}
}
