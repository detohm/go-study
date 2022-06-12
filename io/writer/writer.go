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
	fmt.Println("--- standard lib writer")
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

	// --- custom writer ---
	fmt.Println("--- custom channel writer")
	cWriter := NewChanWriter()
	go func() {
		defer cWriter.Close()
		cWriter.Write([]byte("stream data"))
		cWriter.Write([]byte("this is a book"))
	}()

	for c := range cWriter.Chan() {
		fmt.Printf("%c", c)
	}
	fmt.Println()
	// output
	// stream datathis is a book
}

// -- custom writer ---
type chanWriter struct {
	ch chan byte
}

func NewChanWriter() *chanWriter {
	return &chanWriter{make(chan byte, 1)}
}

func (w *chanWriter) Chan() <-chan byte {
	return w.ch
}

func (w *chanWriter) Write(p []byte) (int, error) {
	n := 0
	for _, b := range p {
		w.ch <- b
		n++
	}
	return n, nil
}

func (w *chanWriter) Close() error {
	close(w.ch)
	return nil
}
