package main

import (
	"fmt"
	"testing"
)

func TestMul(t *testing.T) {
	got := Mul(2, 3)
	expected := 6

	if got != expected {
		t.Errorf("got %v, wanted %v", got, expected)
	}
}

type MulCase struct {
	arg1     int
	arg2     int
	expected int
}

func TestMulTable(t *testing.T) {
	testcases := []MulCase{
		{2, 3, 6},
		{4, 8, 32},
		{10, 20, 200},
	}
	for _, tc := range testcases {
		if got := Mul(tc.arg1, tc.arg2); got != tc.expected {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}

func BenchmarkMul(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mul(2, 3)
	}
}

func ExampleMul() {
	fmt.Println(Mul(2, 3))
	// Output: 6
}

/* Example Commands
go test -v ./test/simple

go test -coverprofile=coverage.out
go tool cover -html=coverage.out

go test -v -bench=./test/simple ./test/simple
go test -v -bench=Mul ./test/simple
*/
