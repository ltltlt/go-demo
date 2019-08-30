package main

import "fmt"

type A struct {
	F1 int `json:"f1"`
	F2 int `json:"f2"`
}
type B struct {
	F1 int `json:"ff1"`
	F2 int `json:"ff2"`
}

func main() {
	a := A{F1: 1, F2: 2}
	b := B(a)
	fmt.Println(b)
}
