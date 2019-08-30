package main

import "github.com/spf13/cast"
import "fmt"

func main() {
	var a interface{} = []string{"1", "2"}
	fmt.Println(cast.ToIntSlice(a)) // cast can cast String slice to int slice(only when all element in the string slice can convert to int), which is awesome

	a = []string{"1", "a"}
	fmt.Println(cast.ToIntSlice(a)) // []

	a = "1"
	fmt.Println(cast.ToInt(a))
}
