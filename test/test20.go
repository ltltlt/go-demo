package main

import (
	"fmt"
)

type A int

const (
	aa = 1
	bb = 2

	a A = iota
	b
	c
	d
)

func main() {
	fmt.Printf("%[1]T %[1]v\n", a)
	fmt.Printf("%[1]T %[1]v\n", b)

	m := make(map[int]int)
	m[1] = m[2] = 3
}
