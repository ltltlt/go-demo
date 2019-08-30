package main

import "fmt"

func main() {
	a, b := AComplexName(8, 0.0)

	a -= 1
	b -= 1
	fmt.Println(a, b)
}

func AComplexName(a int, b float64) (int, float64) {
	for i := 0; i < a; i++ {
		fmt.Println("ok")
	}
	fmt.Println(a, b)

	return 1, 4.0
}
