package main

import (
	"fmt"
	"io"
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
		ch := c.src[c.cur]
		if ch == ' ' {
			ch = '0'
		}
		buf[n] = ch
		n++
		c.cur++
	}
	copy(p, buf)
	return n, nil
}
