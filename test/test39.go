package main

import "fmt"

func test() (i int) {
	v, i := 1, 2
	_ = v
	return
}

func main() {
	fmt.Println(test())
}
