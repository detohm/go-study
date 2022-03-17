package main

import "fmt"

type task struct {
	id    int
	name  string
	point float64
}

// value receiver
func (t task) change1() {
	t.id = 1000
	t.name = "changed1"
}

// pointer receiver
func (t *task) change2() {
	t.id = 2000
	t.name = "changed2"
}

func main() {
	t := task{0, "test1", 34.1234}
	fmt.Printf("before change1 func: %+v\n", t)
	t.change1()
	fmt.Printf("after change1 method called: %+v\n", t)
	t.change2()
	fmt.Printf("after change2 method called: %+v\n", t)
}

/* Result - pointer receiver and value receiver affects behavior
Even though there is no different from calling side.

before change1 func: {id:0 name:test1 point:34.1234}
after change1 method called: {id:0 name:test1 point:34.1234}
after change2 method called: {id:2000 name:changed2 point:34.1234}
*/
