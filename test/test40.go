package main

import "time"
import "fmt"

func test() {
	var start = time.Now()
	defer func() {
		cost := time.Since(start)
		fmt.Println(cost.Nanoseconds())
	}()
}

func main() {
	test()
}
