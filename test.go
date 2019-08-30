package main

import "fmt"

func main() {
	var v = 1
	m := map[*int]int{&v: 2}
	v = 11

	for k, v := range m {
		fmt.Println(*k, v)
	}
}
