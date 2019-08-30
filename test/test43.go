package main

import (
	"fmt"
	"reflect"
)

func main() {
	var v []string

	printKind(v)
}

func printKind(v interface{}) {
	fmt.Println(reflect.TypeOf(v).Kind())
}
