package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

/*
type Reader interface {
	Read(p []byte) (n int, err error)
}
*/

func main() {
	// --- experiment with standard lib reader ---
	/* 	standard library for string reader (strings.Reader)
	type Reader struct {
		s        string
		i        int64 // current reading index
		prevRune int   // index of previous rune; or < 0
	}
	*/
	reader := strings.NewReader("This is very long string.")
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			fmt.Println(string(p[:n]))
			break
		}
		fmt.Println(string(p[:n]))
	}
	/* output (each iteration will output 4-byte string)
	This
	is
	very
	lon
	g st
	ring
	.
	*/

	// --- custom io.Reader ---
	fmt.Println("--- custom reader ---")
	cReader := NewCustomReader("This is very long string.")
	q := make([]byte, 4)
	for {
		n, err := cReader.Read(q)
		if err == io.EOF {
			break
		}
		fmt.Println(string(q[:n]))
	}
	/* output from custom reader
	--- custom reader ---
	This
	0is0
	very
	0lon
	g0st
	ring
	.
	*/

	// --- custom chain io.Reader ---
	fmt.Println("--- custom chain reader ---")
	ccReader := NewCustomChainReader(strings.NewReader("This is an another string."))
	k := make([]byte, 4)
	for {
		n, err := ccReader.Read(k)
		if err == io.EOF {
			break
		}
		fmt.Println(string(k[:n]))
	}
	fmt.Println()
	/* output from custom chain reader
	--- custom chain reader ---
	This
	0is0
	an0a
	noth
	er0s
	trin
	g.
	*/

	// --- custom chain io.Reader benefit with File ---
	/* as you can see, it's very useful as we can read
	from any reader implementation
	*/
	fmt.Println("--- custom chain reader with File ---")
	file, err := os.Open("./io/reader/test-reader.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	cfReader := NewCustomChainReader(file)
	m := make([]byte, 4)
	for {
		n, err := cfReader.Read(m)
		if err == io.EOF {
			break
		}
		fmt.Println(string(m[:n]))
	}
	fmt.Println()
	/* output
	--- custom chain reader with File ---
	This
	0is0
	a0fi
	le!0
	bla0
	bla0
	bla0
	0Ple
	ase0
	read
	0me!
	*/

}

// --- custom io.Reader ---
type customReader struct {
	src string
	cur int
}

func NewCustomReader(src string) *customReader {
	return &customReader{
		src: src,
		cur: 0,
	}
}

func (c *customReader) Read(p []byte) (int, error) {
	if c.cur >= len(c.src) {
		return 0, io.EOF
	}

	dif := len(c.src) - c.cur
	n, bound := 0, 0
	if dif >= len(p) {
		bound = len(p)
	} else if dif <= len(p) {
		bound = dif
	}

	buf := make([]byte, bound)
	for n < bound {
		// put custom logic here
		buf[n] = insertZero(c.src[c.cur])
		n++
		c.cur++
	}
	copy(p, buf)
	return n, nil
}

// --- chain io.Reader ---
type customChainReader struct {
	reader io.Reader
}

func NewCustomChainReader(reader io.Reader) *customChainReader {
	return &customChainReader{
		reader: reader,
	}
}

func (c *customChainReader) Read(p []byte) (int, error) {
	n, err := c.reader.Read(p)
	if err != nil {
		return n, err
	}
	// please note that it might not the same as len(p)
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = insertZero(p[i])
	}

	// copy from buf back to p
	copy(p, buf)
	return n, nil
}

func insertZero(c byte) byte {
	if c == ' ' || c == '\n' || c == '\r' || c == '\t' {
		return '0'
	}
	return c
}
