package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./created-file.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		word, err := reader.ReadString(' ')
		if err != nil {
			if err == io.EOF {
				fmt.Print(word)
				break
			} else {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		fmt.Print(word)
	}
}
