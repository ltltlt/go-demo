package main

import "fmt"
import "unsafe"

type A struct {
	a int32
	b int64
}

func main() {
	a := A{}
	fmt.Println(unsafe.Sizeof(a))
}
