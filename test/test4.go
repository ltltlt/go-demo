package main

import (
	"fmt"
	"github.com/spf13/cast"
)

func main() {
	v := 1.1
	s := fmt.Sprintf("%f", v)
	bs := []byte(s)
	fmt.Println(cast.ToFloat64(bs))
}
