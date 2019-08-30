package main

import "fmt"

type Base struct{}

type Child struct {
	Base
}

func (*Base) test() {
	fmt.Println("base")
}

func (*Child) test() {
	fmt.Println("child")
}

func main() {
	c := Child{}
	c.test()
}
