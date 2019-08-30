package main

import "fmt"

func test() (err error) {
	a, err := 1, fmt.Errorf("err")
	fmt.Println(a, err)
	return
}

func main() {
	fmt.Println(test())
}
