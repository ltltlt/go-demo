package main

import "fmt"

func f(a int) int {
	return a + 1
}

func main() {
	_ = f(4)

	var ff = func(a int) int {
		return a - 1
	}
	_ = ff(3)

	fmt.Println("en")
}
