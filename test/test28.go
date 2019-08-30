package main

import "fmt"

type I interface {
	test()
}

type S struct {
}

func (s *S) test() {
	fmt.Println("s test")
}

func main() {
	var i I = &S{}
	i.test()
	i = nil
	i.test()
}
