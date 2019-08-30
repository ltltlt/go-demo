package main

import "fmt"
import "encoding/json"

type Test struct {
	T []byte
}

func main() {
	t := Test{
		T: []byte{1, 2, 3},
	}
	v, _ := json.Marshal(t)
	fmt.Println(string(v))
}
