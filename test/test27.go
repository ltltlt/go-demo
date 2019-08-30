package main

import "fmt"

type test struct {
	v int
}

func (t *test) print() {
	fmt.Println(t.v)
}

func (t *test) usePrint() {
	t.print()
}

type test1 struct {
	test
}

func (t *test1) print() {
	fmt.Println(t.v + 1)
}

func main() {
	t := &test1{}
	t.usePrint()
}
