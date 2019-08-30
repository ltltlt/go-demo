package main

import "fmt"

func main() {
	var a interface{} = 1
	b, ok := a.(int)
	fmt.Println(b, ok)
}
