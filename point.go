package main

import (
	"fmt"
)

func main() {
	var a [100]struct{}
	var b [100]struct{}
	var v struct{}
	fmt.Printf("%p %p %p\n", &a, &b, &v)
}
