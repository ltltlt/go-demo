package main

import "fmt"

func executeFn(fn func() int) int {
	return fn()
}

func main() {
	a := 1
	b := 2
	go executeFn(func() int {
		a += b
		return a
	})
	fmt.Println(a, b)
}
