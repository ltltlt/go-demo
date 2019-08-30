package main

import (
	"encoding/json"
	"fmt"
)

type test struct {
	a int `json:"a"`
	B int
}

func main() {
	v := &test{a: 1, B: 1}
	bs, err := json.Marshal(v)
	fmt.Println(string(bs), err)
}
