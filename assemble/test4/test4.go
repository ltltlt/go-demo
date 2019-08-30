package main

import (
	"fmt"
	"unsafe"
)

type MyType int

func (a MyType) printthis() {
	print(a)
}

type funcval struct {
	fn uintptr
	a  int
}

func main() {
	var v MyType = 1

	m := v.printthis

	v1 := (**funcval)(unsafe.Pointer(&m))
	fmt.Println(**v1)

	localv := 16

	f := func() {
		println(v)
		println(localv)
	}

	v2 := (**funcval)(unsafe.Pointer(&f))
	fmt.Println(**v2)
}
