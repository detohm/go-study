package main

import "fmt"

type task struct {
	id    int
	name  string
	point float64
}

type ITask interface {
	change1()
	change2()
}

type ITask2 interface {
	change3()
	change4()
}

func main() {

	var t1 ITask
	t1 = task{1, "test1", 34.1234}

	fmt.Printf("t1 before proceed1 func: %+v\n", t1)
	proceed1(t1)
	fmt.Printf("t1 after proceed1 func: %+v\n", t1)

	var t2 ITask2
	t2 = &task{2, "test2", 34.1234}
	// if remove & above it will cause error
	// cannot use (task literal) (value of type task) as ITask2
	// value in assignment: missing method change3
	// (change3 has pointer receiver)

	fmt.Printf("t2 before proceed2 func: %+v\n", t2)
	proceed2(t2)
	fmt.Printf("t2 after proceed2 func: %+v\n", t2)
}

func proceed1(t ITask) {
	t.change1()
	fmt.Printf("after change1 func: %+v\n", t)
}

func proceed2(t ITask2) {
	t.change3()
	fmt.Printf("after change3 func: %+v\n", t)
}

func (t task) change1() {
	t.id = 1000
	t.name = "changed1"
}

func (t task) change2() {
	t.id = 2000
	t.name = "change2"
}

func (t *task) change3() {
	t.id = 3000
	t.name = "change3"
}

func (t *task) change4() {
	t.id = 4000
	t.name = "change4"
}

/* Result
interace contains method signature but implementation can be
either value or pointer sementic
*/
