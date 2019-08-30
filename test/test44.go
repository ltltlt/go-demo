package main

import (
	"fmt"
)

type Structure struct {
	A interface{}
}

func main() {
	v := Structure{[]*Structure{{1}}}
	v1 := v
	v1.A.([]*Structure)[0].A = 2
	fmt.Println(v.A.([]*Structure)[0])
}
