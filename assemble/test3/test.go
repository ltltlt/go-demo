package main

import "fmt"

func main() {
	aSimpleTest()
	bSimpleTest()
	cSimpleTest()
}

func aSimpleTest(){
	fmt.Println("test")
}

func bSimpleTest() {
	println("bbbbbb")
}

func cSimpleTest() {
	println(1, []int{1, 2, 3}, "bbb", 1.1)
}