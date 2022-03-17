package main

import (
	"fmt"
)

type errorCustom struct {
	str  string
	user string
}

func (e *errorCustom) Error() string {
	return fmt.Sprintf("Err: %s from %s", e.str, e.user)
}

func NewErrorCustom(err string, user string) error { // return as interface
	return &errorCustom{err, user}
}

func proceed(i int) error {
	if i%2 == 0 {
		return NewErrorCustom(
			fmt.Sprintf("%d is even", i),
			"mock")
	}
	return nil
}

func main() {
	for i := 0; i < 10; i++ {
		if err := proceed(i); err != nil {
			fmt.Println(err)
		}
	}
}

/* Result

Err: 0 is even from mock
Err: 2 is even from mock
Err: 4 is even from mock
Err: 6 is even from mock
Err: 8 is even from mock

*/
