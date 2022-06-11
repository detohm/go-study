package main

import (
	"bytes"
	"fmt"
	"os"
)

/*
type Writer interface {
	Write(p []byte) (n int, error)
}
*/

func main() {
	// --- experiment with standard lib writer (bytes) ---
	data := []string{
		"Thailand",
		"USA",
		"China",
		"India",
	}
	var writer bytes.Buffer
	for _, line := range data {
		n, err := writer.Write([]byte(line))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if n != len(line) {
			fmt.Println("failed to write data")
			os.Exit(1)
		}
	}
	fmt.Println(writer.String())
	// output
	// ThailandUSAChinaIndia

}
