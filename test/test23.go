package main

import "fmt"

func main() {
	m := map[int]int{
		1:  1,
		2:  2,
		3:  3,
		-1: -1,
	}
	for i, j := range m {
		fmt.Println(i, j)
	}
}
