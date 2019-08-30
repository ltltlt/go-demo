package lib

import "fmt"
import _ "unsafe"

//go:linkname test basic/linkname/lib.Test
func test() {
	fmt.Println("hello world")
}

func Test()
