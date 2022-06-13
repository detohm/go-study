package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	data := []string{
		"This ",
		"is ",
		"a ",
		"book.",
	}
	fmt.Println("start writing file")
	file, err := os.Create("./created-file.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, p := range data {
		n, err := file.Write([]byte(p))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if n != len(p) {
			fmt.Println("failed to write data.")
			os.Exit(1)
		}
	}
	fmt.Println("file write done")
	file.Close()

	fmt.Println("start reading file")
	file2, err := os.Open("./created-file.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file2.Close()

	p := make([]byte, 4)
	for {
		n, err := file2.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Print(string(p[:n]))
	}
	fmt.Println("\nread file done")

}

/* output

start writing file
file write done
start reading file
This is a book.
read file done

*/
