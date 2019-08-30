package main

import "fmt"
import "encoding/json"

type test struct {
	A int
}

func main() {
	t := test{1}
	b, _ := json.Marshal(t)
	json.Unmarshal(b, &t)
	fmt.Println(t)
}
