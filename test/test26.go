package main

import "fmt"

func test() (i int) {
	i = 1
	defer func() {
		i += 100
	}()
	i++
	return i
}

func test1() int {
	i := 1
	defer func() {
		i += 100
	}()
	i++
	return i
}

func main() {
	fmt.Println(test())
	fmt.Println(test1())
}
