package main

import "fmt"

type T struct {
	a int
	b string
}

func (T) A() {}

func main() {
	t := T{1, "a"}
	t1 := struct {
		a int
		b string
	}{
		a: 1,
		b: "a",
	}

	fmt.Println(t == t1)
}
