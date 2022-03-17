package main

import "fmt"

type task struct {
	id    int
	name  string
	point float64
}

func main() {
	t := task{0, "test1", 34.1234}
	fmt.Printf("before change1 func: %+v\n", t)
	change1(t)
	fmt.Printf("after change1 func: %+v\n", t)
	change2(&t)
	fmt.Printf("after change2 func: %+v\n", t)

}

func change1(t task) {
	t.id = 1000
	t.name = "changed1"
}

func change2(t *task) {
	t.id = 2000
	t.name = "change2"
}

/* Result - struct are passed by value in normal basis

before change1 func: {id:0 name:test1 point:34.1234}
after change1 func: {id:0 name:test1 point:34.1234}
after change2 func: {id:2000 name:change2 point:34.1234}
*/
