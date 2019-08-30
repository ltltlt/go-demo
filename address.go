package main

import (
	"fmt"
	"unsafe"
)

type S struct {
	i int
}

func comp(s *S) {
	for i := 0; i < 10; i++ {
		s.i -= i
	}
}

func test() *S {
	s := &S{22}
	for i := 0; i < 10; i++ {
		s.i += i
	}
	comp(s)
	return s
}

var i = 1
var str = "11"

func main() {
	s := &S{2}
	fmt.Printf("%+v\n", s)
	p := unsafe.Pointer(s)
	fmt.Printf("p is: %p\n", p)
	fmt.Printf("p contains: %+v\n", (*S)(p))
	(*S)(p).i = 3
	p1 := unsafe.Pointer(s)
	fmt.Printf("p1 is: %p\n", p1)
	fmt.Printf("p1 contains: %+v\n", (*S)(p1))

	fmt.Printf("test is: %p\n", unsafe.Pointer(test()))
	fmt.Printf("i store in: %p\n", unsafe.Pointer(&i))
	fmt.Printf("str store in: %p\n", unsafe.Pointer(&str))

	var j = 1
	fmt.Printf("j is: %p\n", unsafe.Pointer(&j))
}
