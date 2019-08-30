package main

import "fmt"
import "github.com/spf13/cast"

func main() {
	s := []byte("123.2")
	fmt.Println(cast.ToFloat64(cast.ToString(s)))
}
