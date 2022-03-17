package main

import "fmt"

type task struct {
	id    int
	name  string
	point float64
}

type ITask interface {
	change1()
}

func main() {

	var t1 ITask
	t1 = task{1, "test1", 34.1234}
	fmt.Printf("in main addr: %p\n", &t1)
	proceed1(t1)
}

func proceed1(t ITask) {
	t.change1()
	fmt.Printf("addr in proceed1: %p\n", &t)
	fmt.Printf("after change1 func: %+v\n", t)
}

func (t task) change1() {
	t.id = 1000
	t.name = "change1"
}

/* Result - address of interface is different

in main addr: 0xc000096210
addr in proceed1: 0xc000096220
after change1 func: {id:1 name:test1 point:34.1234}
*/
