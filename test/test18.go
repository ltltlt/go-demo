package main

import "fmt"
import "github.com/spf13/cast"

func main() {
	m := map[string]interface{}{}
	s := cast.ToString(m["1"])
	fmt.Println(len(s))
}
