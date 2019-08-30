package main

import "fmt"

type T int

const (
	a T = iota
	b
	c
	d
)

func main() {
	fmt.Println(a, b, c, d)
}
