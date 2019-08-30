package main

import (
	"fmt"
)

type MyInt int

func (i MyInt) Method() {
	fmt.Println("MyInt method")
}

type MethodIface interface {
	Method()
}

func main() {
	var v MyInt = 1

	var i1 interface{} = v
	var i2 MethodIface = v

	println(i1)
	println(i2)
}
