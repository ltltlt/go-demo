package main

import "unsafe"
import "fmt"

func main() {
	var v [3]uint32
	var vv [3]uint32
	p1 := uintptr(unsafe.Pointer(&v))
	p2 := uintptr(unsafe.Pointer(&vv))
	fmt.Println(p1)
	fmt.Println(p2)
}
