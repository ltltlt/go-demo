package main

import (
	"fmt"
)

func test() float64 {
	return 1.0
}
func main() {
	var value interface{} = 1
	v := (value.(int)) / 100.0
	fmt.Printf("%[1]T %[1]v\n", v)

	v1 := float64(value.(int)) / 100
	fmt.Printf("%[1]T %[1]v\n", v1)

	v2 := test() / 100
	fmt.Printf("%[1]T %[1]v\n", v2)
}
