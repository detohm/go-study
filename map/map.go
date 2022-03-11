package main

import "fmt"

func main() {

	var m map[string]int
	// m["first"] = 0 - it will cause error assignment to entry in nil map
	fmt.Printf("m:%+v\n", m)

	m2 := make(map[string]int)
	m2["zero"] = 0
	fmt.Printf("m2:%v\n", m2)
	// Result: m2:map[zero:0]
	m3 := m2
	m3["one"] = 1
	fmt.Printf("m3:%v m2:%v\n", m3, m2)
	// Result:
	// m3:map[one:1 zero:0] m2:map[one:1 zero:0]
	// side effect on m2 if add element into m3

	m4 := make(map[int]int)
	m5 := m4
	for i := 0; i < 100001; i++ {
		m4[i] = i * 2
	}

	fmt.Printf("m5[100000]:%d m4[100000]:%d m5-len:%d m4-len:%d\n",
		m5[100000], m4[100000],
		len(m5), len(m4))
	// Result:
	// m5[100000]:200000 m4[100000]:200000 m5-len:100001 m4-len:100001
	// Note: There is no dynamic allocation
	// and no change in underlying data structure
	// (not the same machanism as slice)

	m6 := make(map[string]int)
	update(m6)
	fmt.Printf("m6:%+v\n", m6)
	// Result:
	// m6:map[add-in-func:1]
	// Note: even passing by value, the result is affected
	// as it's the reference type

}

func update(m map[string]int) {
	m["add-in-func"] = 1
}
