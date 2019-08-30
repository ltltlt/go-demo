package main

import "fmt"

func main() {
	var values []interface{}
	values1 := []int{1, 2, 3, 4, 5}

	for _, v := range values1 {
		values = append(values, &v)
	}
	fmt.Println(values)
}
