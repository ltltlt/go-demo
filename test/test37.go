package main

import "fmt"

type T struct {
	a int
}

func (T) A() {
	fmt.Println("abc")
}

func main() {
	var t *T
	t.A()
}
