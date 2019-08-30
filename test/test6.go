package main

import "fmt"

func test() (a int) {
	a, b := 1, 2
	fmt.Println(b)
	return
}

func main() {
	fmt.Println(test())
}
