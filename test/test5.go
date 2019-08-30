package main

import "fmt"

func test() (err error) {
	err = fmt.Errorf("no")
	defer func() {
		if err != nil {
			return
		}
		if e := recover(); e != nil {
			return
		}
	}()
	panic("what")
	return
}
func main() {
	t := test()
	fmt.Printf("%[1]T %[1]v", t)
}
