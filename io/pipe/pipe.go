package main

import (
	"bytes"
	"io"
	"os"
	"time"
)

func main() {

	buf := new(bytes.Buffer)
	buf.WriteString("This is a book\n")

	reader, writer := io.Pipe()

	go func() {
		time.Sleep(time.Second * 2)
		defer writer.Close()
		io.Copy(writer, buf)

	}()
	// copy got blocked from reader
	io.Copy(os.Stdout, reader)
	reader.Close()
}

/* output
program got blocked for 2 seconds
before output: This is a book
*/
