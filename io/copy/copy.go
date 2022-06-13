package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	// io.Copy - arbitrary function to stream data from
	// source reader to target writer.

	data := new(bytes.Buffer)
	data.WriteString("this ")
	data.WriteString("is ")
	data.WriteString("a ")
	data.WriteString("book from copy.")

	file, err := os.Create("./created-file.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	if _, err := io.Copy(file, data); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("successfully copied from buffer to file")

	file2, err := os.Open("./created-file.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file2.Close()

	if _, err := io.Copy(os.Stdout, file2); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("\nsuccessfully copied from file to stdout")
}

/* output
successfully copied from buffer to file
this is a book from copy.
successfully copied from file to stdout
*/
