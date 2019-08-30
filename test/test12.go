package main

import "fmt"

func main() {
	type test struct {
		a int
	}
	values := []test{
		test{1},
		test{2},
		test{3},
	}
	for _, v := range values {
		v.a = 4
	}
	fmt.Println(values)
}
